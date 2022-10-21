package week02

import (
	"errors"
	"reflect"
	"strings"
)

var errInvalidEntity = errors.New("invalid entity")

// InsertStatement 通过反射将结构体转换成 insert sql
func InsertStatement(entity interface{}) (string, []interface{}, error) {
	if nil == entity {
		return "", nil, errInvalidEntity
	}

	t4type := reflect.TypeOf(entity)
	t4val := reflect.ValueOf(entity)

	// 处理一级结构体指针
	if reflect.Pointer == t4type.Kind() {
		t4type = t4type.Elem()
		t4val = t4val.Elem()
	}

	if t4type.Kind() != reflect.Struct {
		return "", nil, errInvalidEntity
	}

	t4num := t4type.NumField()
	if t4num <= 0 {
		return "", nil, errInvalidEntity
	}

	s5field, s5val, err := structureAnalysisForInsert(t4val.Interface())
	if err != nil {
		return "", nil, err
	}

	sql := makeInsertSql(t4type.Name(), s5field)

	return sql, s5val, nil
}

// structureAnalysisForInsert 结构分析
func structureAnalysisForInsert(entity interface{}) ([]string, []interface{}, error) {
	t4type := reflect.TypeOf(entity)
	t4val := reflect.ValueOf(entity)

	t4num := t4type.NumField()
	if t4num <= 0 {
		return []string{}, []interface{}{}, nil
	}

	s5field := make([]string, 0, t4num)
	s5val := make([]interface{}, 0, t4num)
	m3key := make(map[string]bool, t4num)
	// 遍历属性
	for i := 0; i < t4num; i++ {
		t4field := t4type.Field(i)
		t4fv := t4val.Field(i)

		t4fvk := t4fv.Kind()
		//if reflect.Pointer == t4fvk {
		//	// 这里有可能有方法指针
		//	if !t4fv.IsNil() && reflect.Struct == t4fv.Elem().Kind() {
		//		t4fv = t4fv.Elem()
		//	}
		//}

		t4fvpp := t4fv.Type().PkgPath()
		if "database/sql" == t4fvpp {
			switch t4fv.Type().Name() {
			case "NullString", "NullInt32":
				t4fv = t4val.Field(i)
				m3key[t4field.Name] = true
				s5field = append(s5field, t4field.Name)
				if t4field.IsExported() {
					s5val = append(s5val, t4val.Field(i).Interface())
				} else {
					s5val = append(s5val, reflect.Zero(t4field.Type).Interface())
				}
			}

			continue
		}

		if reflect.Struct == t4fvk && t4field.Anonymous {
			t4s5field, t4s5val, t4err := structureAnalysisForInsert(t4fv.Interface())
			if nil != t4err {
				return nil, nil, t4err
			}
			for m := 0; m < len(t4s5field); m++ {
				if _, ok := m3key[t4s5field[m]]; ok {
					continue
				}
				m3key[t4s5field[m]] = true
				s5field = append(s5field, t4s5field[m])
				s5val = append(s5val, t4s5val[m])
			}
			continue
		}

		m3key[t4field.Name] = true
		s5field = append(s5field, t4field.Name)
		if t4field.IsExported() {
			s5val = append(s5val, t4fv.Interface())
		} else {
			s5val = append(s5val, reflect.Zero(t4field.Type).Interface())
		}
	}

	return s5field, s5val, nil
}

// 构造 insert sql
func makeInsertSql(name string, s5field []string) string {
	sb := strings.Builder{}

	sb.WriteString("INSERT INTO `")
	sb.WriteString(name)

	sb.WriteString("`(`")
	for i := 0; i < len(s5field)-1; i++ {
		sb.WriteString(s5field[i])
		sb.WriteString("`,`")
	}
	sb.WriteString(s5field[len(s5field)-1])
	sb.WriteString("`)")

	sb.WriteString(" VALUES")

	sb.WriteString("(")
	for i := 0; i < len(s5field)-1; i++ {
		sb.WriteString("?,")
	}
	sb.WriteString("?);")

	return sb.String()
}
