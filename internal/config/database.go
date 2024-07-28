package config

import "xorm.io/xorm"

var engine *xorm.Engine

func Init() {
	//var err error
	//engine, err = xorm.NewEngine("mysql", "root:12345678@tcp(127.0.0.1:3306)/ems-admin-go?charset=utf8")
	//if err != nil {
	//	panic(err)
	//}
}
