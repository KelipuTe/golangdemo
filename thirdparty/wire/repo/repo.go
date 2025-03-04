package repo

import "demo-golang/thirdparty/wire/repo/dao"

type Repo struct {
	dao *dao.Dao
}

func NewRepo(dao *dao.Dao) *Repo {
	return &Repo{dao}
}
