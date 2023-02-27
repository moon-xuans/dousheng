package user_info

import (
	"dousheng/cache"
	"dousheng/model"
)

type FollowerList struct {
	UserList []*model.UserInfo `json:"user_list"`
}

func QueryFollowerList(userId int64) (*FollowerList, error) {
	return NewQueryFollowerListFlow(userId).Do()
}

type QueryFollowerListFlow struct {
	userId int64

	userList []*model.UserInfo

	*FollowerList
}

func NewQueryFollowerListFlow(userId int64) *QueryFollowerListFlow {
	return &QueryFollowerListFlow{userId: userId}
}

func (q *QueryFollowerListFlow) Do() (*FollowerList, error) {
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

	return q.FollowerList, nil
}

func (q *QueryFollowerListFlow) checkData() error {
	if !model.NewUserInfoDAO().IsUserExistById(q.userId) {
		return ErrUserNotExist
	}
	return nil
}

func (q *QueryFollowerListFlow) prepareData() error {
	err := model.NewUserInfoDAO().GetFollowerListByUserId(q.userId, &q.userList)
	if err != nil {
		return err
	}
	for _, v := range q.userList {
		v.IsFollow = cache.NewProxyIndexMap().GetUserRelation(q.userId, v.Id)
		//v.IsFollow = cache.NewProxyIndexMap().GetUserRelation(v.Id, q.userId) // I think
	}
	return nil
}

func (q *QueryFollowerListFlow) packData() error {
	q.FollowerList = &FollowerList{UserList: q.userList}
	return nil
}
