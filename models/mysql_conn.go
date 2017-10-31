package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	_OB_MYSQL_CONN = "bdm286382118:czs0510016net@tcp(bdm286382118.my3w.com)/bdm286382118_db?charset=utf8"
	_MYSQL_DDRIVER = "mysql"
)

type Movie struct {
	Id int64
}

func RegisterDB() {
	//orm.RegisterModel(new(Movie))
	//orm.RegisterDriver(_MYSQL_DDRIVER, orm.DRMySQL)
	orm.RegisterDataBase("default", _MYSQL_DDRIVER, _OB_MYSQL_CONN, 50)
}
