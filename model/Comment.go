package model

import "time"

type Comment struct {
	ID         int64     `json:"id"`            // 评论id
	UserInfoId int64     `json:"-"`             // 用于一对多关系的id
	VideoId    int64     `json:"-"`             // 一对多,视频对评论
	User       UserInfo  `json:"user" gorm:"-"` // 评论用户信息
	Content    string    `json:"content"`       // 评论内容
	CreatedAt  time.Time `json:"-"`
	CreateDate string    `json:"create_date" gorm:"-"` // 评论发布日期，格式 mm-dd
}
