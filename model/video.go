package model

import (
	"errors"
	"sync"
	"time"
)

type Video struct {
	Id            int64       `json:"id,omitempty"`
	UserInfoId    int64       `json:"-"`
	Author        UserInfo    `json:"author,omitempty" gorm:"-"` // 这里是作者对视频的一对多的关系，而不是视频对作者，故gorm不能存他，但json需要返回他
	PlayUrl       string      `json:"play_url,omitempty"`        // 视频播放地址
	CoverUrl      string      `json:"cover_url,omitempty"`       // 视频封面地址
	FavoriteCount int64       `json:"favorite_count,omitempty"`  // 视频的点赞总数
	CommentCount  int64       `json:"comment_count,omitempty"`   // 视频的评论总数
	IsFavorite    bool        `json:"is_favorite,omitempty"`     // true-已点赞，false-未点赞
	Title         string      `json:"title,omitempty"`           // 视频标题
	Users         []*UserInfo `json:"-" gorm:"many2many:user_favor_videos;"`
	Comments      []*Comment  `json:"-"`
	CreateAt      time.Time   `json:"-"`
	UpdateAt      time.Time   `json:"-"`
}

type VideoDAO struct {
}

var (
	videoDAO  *VideoDAO
	videoOnce sync.Once
)

func NewVideoDAO() *VideoDAO {
	videoOnce.Do(func() {
		videoDAO = new(VideoDAO)
	})
	return videoDAO
}

// AddVideo 添加视频
// 注意：由于视频和userInfo存在多对一的关系，所以传入的Video参数一定要进行id的映射处理
func (v *VideoDAO) AddVideo(video *Video) error {
	if video == nil {
		return errors.New("AddVideo video 空指针")
	}
	return DB.Create(video).Error
}

func (v *VideoDAO) QueryVideoCountByUserId(userId int64, count *int64) error {
	if count == nil {
		return errors.New("QueryVideoCountByUserId count 空指针")
	}
	return DB.Model(&Video{}).Where("user_info_id = ?", userId).Count(count).Error
}
