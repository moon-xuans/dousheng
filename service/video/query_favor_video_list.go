package video

import (
	"dousheng/model"
	"errors"
)

type FavorList struct {
	Videos []*model.Video `json:"video_list"`
}

func QueryFavorVideoList(userId int64) (*FavorList, error) {
	return NewQueryFavorVideoListFlow(userId).Do()
}

type QueryFavorListFlow struct {
	userId int64

	videos []*model.Video

	videoList *FavorList
}

func NewQueryFavorVideoListFlow(userId int64) *QueryFavorListFlow {
	return &QueryFavorListFlow{userId: userId}
}

func (q *QueryFavorListFlow) Do() (*FavorList, error) {
	if err := q.checkData(); err != nil {
		return nil, err
	}
	if err := q.prepareData(); err != nil {
		return nil, err
	}
	if err := q.packData(); err != nil {
		return nil, err
	}
	return q.videoList, nil
}

func (q *QueryFavorListFlow) checkData() error {
	if !model.NewUserInfoDAO().IsUserExistById(q.userId) {
		return errors.New("用户状态异常")
	}
	return nil
}

func (q *QueryFavorListFlow) prepareData() error {
	err := model.NewVideoDAO().QueryFavorVideoListByUserId(q.userId, &q.videos)
	if err != nil {
		return err
	}
	// 填充信息(Author和IsFavorite字段，由于是点赞列表，故所有的都是点赞状态)
	for i := range q.videos {
		// 作者信息查询
		var userInfo model.UserInfo
		err = model.NewUserInfoDAO().QueryUserInfoById(q.videos[i].UserInfoId, &userInfo)
		if err != nil { // 若查询未出错则更新，否则不更新作者信息
			q.videos[i].Author = userInfo
		}
		q.videos[i].IsFavorite = true
	}
	return nil
}

func (q *QueryFavorListFlow) packData() error {
	q.videoList = &FavorList{Videos: q.videos}
	return nil
}
