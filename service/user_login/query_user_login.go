package user_login

import "dousheng/model"

type LoginResponse struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

func QueryUserLogin(username, password string) (*LoginResponse, error) {
	return NewQueryUserLoginFlow(username, password).Do()
}

func NewQueryUserLoginFlow(username, password string) *QueryUserLoginFlow {
	return &QueryUserLoginFlow{username: username, password: password}
}

type QueryUserLoginFlow struct {
	username string
	password string

	data   *LoginResponse
	userId int64
	token  string
}

func (q *QueryUserLoginFlow) Do() (*LoginResponse, error) {
	// 对参数进行合法性验证
	// 准备好数据
	if err := q.prepareData(); err != nil {
		return nil, err
	}
	// 打包最终数据
	if err := q.packData(); err != nil {
		return nil, err
	}
	return q.data, nil
}

func (q *QueryUserLoginFlow) prepareData() error {
	userLoginDAO := model.NewUserLoginDao()
	var login model.UserLogin
	// 准备好userId
	err := userLoginDAO.QueryUserLogin(q.username, q.password, &login)
	if err != nil {
		return err
	}
	q.userId = login.UserInfoId

	// token
	token := q.username + q.password

	q.token = token
	return nil
}

func (q *QueryUserLoginFlow) packData() error {
	q.data = &LoginResponse{
		UserId: q.userId,
		Token:  q.token,
	}
	return nil
}
