package model

import (
	"errors"
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

func (u *UserLoginDAO) QueryUserLogin(username string, password string, login *UserLogin) error {
	if login == nil {
		return errors.New("结构体指针为空")
	}
	DB.Where("username = ? and password = ?", username, password).First(login)
	if login.Id == 0 {
		return errors.New("用户不存在，账号或密码错误")
	}
	return nil
}

func (u *UserLoginDAO) IsUserExistByUsername(username string) bool {
	var userLogin UserLogin
	DB.Where("username = ?", username).First(&userLogin)
	if userLogin.Id == 0 {
		return false
	}
	return true
}
