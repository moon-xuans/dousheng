package comment

import (
	"dousheng/model"
	"dousheng/util"
	"errors"
	"fmt"
)

const (
	CREATE = 1
	DELETE = 2
)

type Response struct {
	Comment *model.Comment `json:"comment"`
}

func PostComment(userId, videoId, commentId, actionType int64, commentText string) (*Response, error) {
	return NewPostCommentFlow(userId, videoId, commentId, actionType, commentText).Do()
}

type PostCommentFlow struct {
	userId      int64
	videoId     int64
	commentId   int64
	actionType  int64
	commentText string

	comment *model.Comment

	*Response
}

func NewPostCommentFlow(userId int64, videoId, commentId, actionType int64, commentText string) *PostCommentFlow {
	return &PostCommentFlow{userId: userId, videoId: videoId, commentId: commentId, actionType: actionType, commentText: commentText}
}

func (p *PostCommentFlow) Do() (*Response, error) {
	var err error
	if err = p.checkData(); err != nil {
		return nil, err
	}
	if err = p.prepareData(); err != nil {
		return nil, err
	}
	if err = p.packData(); err != nil {
		return nil, err
	}
	return p.Response, nil
}

// CreateComment 增加评论
func (p *PostCommentFlow) CreateComment() (*model.Comment, error) {
	comment := model.Comment{UserInfoId: p.userId, VideoId: p.videoId, Content: p.commentText}
	err := model.NewCommentDAO().AddCommentAndUpdateCount(&comment)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (p *PostCommentFlow) DeleteComment() (*model.Comment, error) {
	// 获取comment
	var comment model.Comment
	err := model.NewCommentDAO().QueryCommentById(p.commentId, &comment)
	if err != nil {
		return nil, err
	}
	// 删除comment
	err = model.NewCommentDAO().DeleteCommentAndUpdateCountById(p.commentId, p.videoId)
	return &comment, nil
}

func (p *PostCommentFlow) checkData() error {
	if !model.NewUserInfoDAO().IsUserExistById(p.userId) {
		return fmt.Errorf("用户%d不存在", p.userId)
	}
	if !model.NewVideoDAO().IsVideoExistById(p.videoId) {
		return fmt.Errorf("视频%d不存在", p.videoId)
	}
	if p.actionType != CREATE && p.actionType != DELETE {
		return errors.New("未定义的行为")
	}
	return nil
}

func (p *PostCommentFlow) prepareData() error {
	var err error
	switch p.actionType {
	case CREATE:
		p.comment, err = p.CreateComment()
	case DELETE:
		p.comment, err = p.DeleteComment()
	default:
		return errors.New("未定义的行为")
	}
	return err
}

func (p *PostCommentFlow) packData() error {
	// 填充字段
	userInfo := model.UserInfo{}
	_ = model.NewUserInfoDAO().QueryUserInfoById(p.comment.UserInfoId, &userInfo)
	p.comment.User = userInfo
	_ = util.FillCommentFields(p.comment)

	p.Response = &Response{Comment: p.comment}
	return nil
}
