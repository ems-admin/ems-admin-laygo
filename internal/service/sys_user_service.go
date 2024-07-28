package service

import (
	"ems-adming-go/internal/config"
	"ems-adming-go/internal/model"
	"fmt"
)

type SysUserService struct {
}

func FindByName(username string) (*model.SysUser, error) {
	var sysUser model.SysUser
	has, err := config.Engine.Table("sys_user").Where("username = ?", username).Get(&sysUser)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, fmt.Errorf("用户不存在")
	}
	return &sysUser, nil
}
