package user_info

import (
	"dousheng/model"
	"errors"
)

var (
	ErrUserNotExist = errors.New("用户不存在或已注销")
)

type FollowList struct {
	UserList []*model.UserInfo `json:"user_list"`
}

func QueryFollowList(userId int64) (*FollowList, error) {
	return NewQueryFollowListFlow(userId).Do()
}

type QueryFollowListFlow struct {
	userId int64

	userList []*model.UserInfo

	*FollowList
}

func NewQueryFollowListFlow(userId int64) *QueryFollowListFlow {
	return &QueryFollowListFlow{userId: userId}
}

func (q *QueryFollowListFlow) Do() (*FollowList, error) {
	var err error
	if err = q.checkData(); err != nil {
		return nil, err
	}
	if err = q.prepareData(); err != nil {
		return nil, err
	}
	if err = q.packData(); err != nil {
		return nil, err
	}

	return q.FollowList, nil
}

func (q *QueryFollowListFlow) checkData() error {
	if !model.NewUserInfoDAO().IsUserExistById(q.userId) {
		return ErrUserNotExist
	}
	return nil
}

func (q *QueryFollowListFlow) prepareData() error {
	var userList []*model.UserInfo
	err := model.NewUserInfoDAO().GetFollowListByUserId(q.userId, &userList)
	if err != nil {
		return err
	}
	for i, _ := range userList {
		userList[i].IsFollow = true // 当前用户的关注列表，故isFollow一定为true
	}
	q.userList = userList
	return nil
}

func (q *QueryFollowListFlow) packData() error {
	q.FollowList = &FollowList{UserList: q.userList}
	return nil
}
