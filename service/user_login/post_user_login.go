package user_login

import (
	"dousheng/models"
	"errors"
)

func PostUserLogin(username, password string) (*LoginResponse, error) {
	return NewPostUserLoginFlow(username, password).Do()
}

func NewPostUserLoginFlow(username, password string) *PostUserLoginFlow {
	return &PostUserLoginFlow{username: username, password: password}
}

type PostUserLoginFlow struct {
	username string
	password string

	data   *LoginResponse
	userId int64
	token  string
}

func (q *PostUserLoginFlow) Do() (*LoginResponse, error) {
	// 对参数进行合法性验证

	// 更新数据到数据库
	if err := q.UpdateData(); err != nil {

	}
	// 打包response

	return q.data, nil
}

func (q *PostUserLoginFlow) UpdateData() error {
	// 准备好userInfo,默认name为username
	userLogin := models.UserLogin{Username: q.username, Password: q.password}
	userInfo := models.UserInfo{User: &userLogin, Name: q.username}

	// 判断用户名是否已经存在
	userLoginDAO := models.NewUserLoginDao()
	if userLoginDAO.IsUserExistByUsername(q.username) {
		return errors.New("用户已存在")
	}

	// 更新操作,由于userLogin属于userInfo，故更新userInfo即可，并且传入的是指针
	userInfoDAO := models.NewUserInfoDAO()
	err := userInfoDAO.AddUserInfo(&userInfo)
	if err != nil {
		return err
	}

	q.token = q.username + q.password
	q.userId = userInfo.Id
	return nil
}
