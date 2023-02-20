package user_login

import (
	"dousheng/middleware"
	"dousheng/model"
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
	if err := q.checkData(); err != nil {
		return nil, err
	}
	// 更新数据到数据库
	if err := q.updateData(); err != nil {
		return nil, err
	}
	// 打包response
	if err := q.packResponse(); err != nil {
		return nil, err
	}
	return q.data, nil
}

func (q *PostUserLoginFlow) checkData() error {
	if q.username == "" {
		return errors.New("用户名为空")
	}
	if len(q.username) > MaxUsernameLength {
		return errors.New("用户名长度超出限制")
	}
	if q.password == "" {
		return errors.New("密码为空")
	}
	return nil
}

func (q *PostUserLoginFlow) updateData() error {
	// 准备好userInfo,默认name为username
	userLogin := model.UserLogin{Username: q.username, Password: q.password}
	userInfo := model.UserInfo{User: &userLogin, Name: q.username}

	// 判断用户名是否已经存在
	userLoginDAO := model.NewUserLoginDao()
	if userLoginDAO.IsUserExistByUsername(q.username) {
		return errors.New("用户已存在")
	}

	// 更新操作,由于userLogin属于userInfo，故更新userInfo即可，并且传入的是指针
	userInfoDAO := model.NewUserInfoDAO()
	err := userInfoDAO.AddUserInfo(&userInfo)
	if err != nil {
		return err
	}

	// 颁发token
	token, err := middleware.ReleaseToken(userLogin)
	if err != nil {
		return err
	}
	q.token = token
	q.userId = userInfo.Id
	return nil
}

func (q *PostUserLoginFlow) packResponse() error {
	q.data = &LoginResponse{
		UserId: q.userId,
		Token:  q.token,
	}
	return nil
}
