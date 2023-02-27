package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID         int64     `json:"id"`            // 评论id
	UserInfoId int64     `json:"-"`             // 用于一对多关系的id
	VideoId    int64     `json:"-"`             // 一对多,视频对评论
	User       UserInfo  `json:"user" gorm:"-"` // 评论用户信息
	Content    string    `json:"content"`       // 评论内容
	CreatedAt  time.Time `json:"-"`
	CreateDate string    `json:"create_date" gorm:"-"` // 评论发布日期，格式 mm-dd
}

type CommentDAO struct {
}

var (
	commentDao CommentDAO
)

func NewCommentDAO() *CommentDAO {
	return &commentDao
}

func (c *CommentDAO) AddCommentAndUpdateCount(comment *Comment) error {
	if comment == nil {
		return errors.New("AddCommentAndUpdateCount comment 空指针")
	}
	// 执行事务
	return DB.Transaction(func(tx *gorm.DB) error {
		// 添加评论数据
		if err := tx.Create(comment).Error; err != nil {
			return err
		}
		// 增加count
		if err := tx.Exec("UPDATE videos v SET v.comment_count=v.comment_count+1 WHERE v.id = ?", comment.VideoId).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})
}

func (c *CommentDAO) DeleteCommentAndUpdateCountById(commentId, videoId int64) error {
	// 执行事务
	return DB.Transaction(func(tx *gorm.DB) error {
		// 删除评论
		if err := tx.Exec("DELETE FROM comments WHERE id = ?", commentId).Error; err != nil {
			return err
		}
		if err := tx.Exec("UPDATE videos v SET v.comment_count = v.comment_count - 1 WHERE v.id = ?", videoId).Error; err != nil {
			return nil
		}
		return nil
	})
}

func (c *CommentDAO) QueryCommentById(commentId int64, comment *Comment) error {
	if comment == nil {
		return errors.New("QueryCommentById comment 空指针")
	}
	return DB.Where("id = ?", commentId).First(comment).Error
}

func (c *CommentDAO) QueryCommentListByVideoId(videoId int64, comments *[]*Comment) error {
	if comments == nil {
		return errors.New("QueryCommentListByVideoId comments 空指针")
	}
	if err := DB.Where("video_id = ?", videoId).Find(comments).Error; err != nil {
		return err
	}
	return nil
}
