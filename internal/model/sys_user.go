package model

import "time"

type SysUser struct {
	Id         int64     `xorm:"pk autoincr"` // 主键，自增
	Email      string    `xorm:"varchar(100)"`
	Username   string    `xorm:"varchar(50)"`
	Password   string    `xorm:"varchar(100)"`
	NickName   string    `xorm:"varchar(50)"`
	Enabled    bool      `xorm:"bit(1)"`
	CreateTime time.Time `xorm:"datetime"`
	updateTime time.Time `xorm:"datetime"`
}
