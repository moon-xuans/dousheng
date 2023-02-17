package models

import (
	"sync"
)

type UserLogin struct {
	Id         int64 `gorm:"primary_key"`
	UserInfoId int64
	Username   string `gorm:"primary_key"`
	Password   string `gorm:"size:200;not null"`
}

type UserLoginDAO struct {
}

var (
	userLoginDao  *UserLoginDAO
	userLoginOnce sync.Once
)

func NewUserLoginDao() *UserLoginDAO {
	userLoginOnce.Do(func() {
		userLoginDao = new(UserLoginDAO)
	})
	return userLoginDao
}

func (d UserLoginDAO) IsUserExistByUsername(username string) bool {
	var userLogin UserLogin
	DB.Where("username = ?", username).First(&userLogin)
	if userLogin.Id == 0 {
		return false
	}
	return true
}
