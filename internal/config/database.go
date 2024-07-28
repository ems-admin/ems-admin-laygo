package config

import (
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

func InitDB() error {
	var err error
	Engine, err = xorm.NewEngine("mysql", "root:12345678@tcp(127.0.0.1:3306)/ems-admin-go?charset=utf8")
	if err != nil {
		panic(err)
		return err
	}

	Engine.ShowSQL(true)
	Engine.SetLogLevel(4)

	if err = Engine.Ping(); err != nil {
		panic(err)
		return err
	}
	return nil
}

func CloseDB() {
	if err := Engine.Close(); err != nil {
		panic(err)
	}
}
