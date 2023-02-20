package video

import (
	"dousheng/cache"
	"dousheng/model"
	"errors"
)

type Videos struct {
	Videos []*model.Video `json:"video_list,omitempty"`
}

func QueryVideoListByUserId(userId int64) (*Videos, error) {
	return NewQueryVideoListByUserIdFlow(userId).Do()
}

func NewQueryVideoListByUserIdFlow(userId int64) *QueryVideoListByUserIdFlow {
	return &QueryVideoListByUserIdFlow{userId: userId}
}

type QueryVideoListByUserIdFlow struct {
	userId int64
	videos []*model.Video

	videoList *Videos
}

func (q *QueryVideoListByUserIdFlow) Do() (*Videos, error) {
	if err := q.checkData(); err != nil {
		return nil, err
	}
	if err := q.packData(); err != nil {
		return nil, err
	}
	return q.videoList, nil
}

func (q *QueryVideoListByUserIdFlow) checkData() error {
	// 检查userId是否存在
	if !model.NewUserInfoDAO().IsUserExistById(q.userId) {
		return errors.New("用户不存在")
	}
	return nil
}

// 注意：Video由于在数据库中没有存储作者信息，所以需要手动填充
func (q *QueryVideoListByUserIdFlow) packData() error {
	err := model.NewVideoDAO().QueryVideoListByUserId(q.userId, &q.videos)
	if err != nil {
		return err
	}
	// 作者信息查询
	var userInfo model.UserInfo
	err = model.NewUserInfoDAO().QueryUserInfoById(q.userId, &userInfo)
	p := cache.NewProxyIndexMap()
	if err != nil {
		return err
	}
	// 填充信息(Author和IsFavorite字段)
	for i := range q.videos {
		q.videos[i].Author = userInfo
		q.videos[i].IsFavorite = p.GetVideoFavorState(q.userId, q.videos[i].Id)
	}

	q.videoList = &Videos{Videos: q.videos}

	return nil
}
