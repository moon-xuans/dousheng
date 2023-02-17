package models

import (
	"errors"
	"sync"
)

var (
	ErrIvdPtr        = errors.New("空指针错误")
	ErrEmptyUserList = errors.New("用户列表为空")
)

type UserInfo struct {
	Id   int64      `json:"id" gorm:"id, omitempty"`
	Name string     `json:"name" gorm:"name, omitempty"`
	User *UserLogin `json:"-"`
}

type UserInfoDAO struct {
}

var (
	userInfoDAO  *UserInfoDAO
	userInfoOnce sync.Once
)

func NewUserInfoDAO() *UserInfoDAO {
	userInfoOnce.Do(func() {
		userInfoDAO = new(UserInfoDAO)
	})
	return userInfoDAO
}

func (u *UserInfoDAO) AddUserInfo(userInfo *UserInfo) error {
	if userInfo == nil {
		return ErrIvdPtr
	}
	return DB.Create(userInfo).Error
}
