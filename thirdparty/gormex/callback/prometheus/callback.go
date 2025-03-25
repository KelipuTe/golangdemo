package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"gorm.io/gorm"
	"time"
)

// Callback gorm的普罗米修斯监控插件
// 监控，查询语句执行时间
type Callback struct {
	vector *prometheus.SummaryVec
}

func NewCallback(opts prometheus.SummaryOpts) *Callback {
	vector := prometheus.NewSummaryVec(
		opts,
		[]string{"type", "table"},
	)
	prometheus.MustRegister(vector)
	return &Callback{
		vector: vector,
	}
}

func (t *Callback) Name() string {
	return "prometheus"
}

func (t *Callback) Before() func(db *gorm.DB) {
	return func(db *gorm.DB) {
		start := time.Now()
		db.Set("starttime", start)
	}
}

func (t *Callback) After(queryType string) func(db *gorm.DB) {
	return func(db *gorm.DB) {
		val, _ := db.Get("starttime")
		start, ok := val.(time.Time)
		if ok {
			timeSince := time.Since(start).Milliseconds()
			t.vector.WithLabelValues(queryType, db.Statement.Table).
				Observe(float64(timeSince))
		}
	}
}

func (t *Callback) Initialize(db *gorm.DB) error {
	err := db.Callback().Create().Before("*").
		Register("prometheus_create_before", t.Before())
	if err != nil {
		return err
	}
	err = db.Callback().Create().After("*").
		Register("prometheus_create_after", t.After("CREATE"))
	if err != nil {
		return err
	}

	err = db.Callback().Query().Before("*").
		Register("prometheus_query_before", t.Before())
	if err != nil {
		return err
	}
	err = db.Callback().Query().After("*").
		Register("prometheus_query_after", t.After("QUERY"))
	if err != nil {
		return err
	}

	err = db.Callback().Update().Before("*").
		Register("prometheus_update_before", t.Before())
	if err != nil {
		return err
	}
	err = db.Callback().Update().After("*").
		Register("prometheus_update_after", t.After("UPDATE"))
	if err != nil {
		return err
	}

	err = db.Callback().Delete().Before("*").
		Register("prometheus_delete_before", t.Before())
	if err != nil {
		return err
	}
	err = db.Callback().Delete().After("*").
		Register("prometheus_delete_after", t.After("DELETE"))
	if err != nil {
		return err
	}

	err = db.Callback().Row().Before("*").
		Register("prometheus_row_before", t.Before())
	if err != nil {
		return err
	}
	err = db.Callback().Row().After("*").
		Register("prometheus_row_after", t.After("ROW"))
	if err != nil {
		return err
	}

	err = db.Callback().Raw().Before("*").
		Register("prometheus_raw_before", t.Before())
	if err != nil {
		return err
	}
	err = db.Callback().Raw().After("*").
		Register("prometheus_raw_after", t.After("RAW"))
	if err != nil {
		return err
	}

	return err
}
