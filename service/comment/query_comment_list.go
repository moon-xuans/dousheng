package comment

import (
	"dousheng/model"
	"dousheng/util"
	"fmt"
)

type List struct {
	Comments []*model.Comment `json:"comment_list"`
}

func QueryCommentList(userId, videoId int64) (*List, error) {
	return NewQueryCommentListFlow(userId, videoId).Do()
}

type QueryCommentListFlow struct {
	userId  int64
	videoId int64

	comments []*model.Comment

	commentList *List
}

func NewQueryCommentListFlow(userId, videoId int64) *QueryCommentListFlow {
	return &QueryCommentListFlow{userId: userId, videoId: videoId}
}

func (q *QueryCommentListFlow) Do() (*List, error) {
	if err := q.checkData(); err != nil {
		return nil, err
	}
	if err := q.prepareData(); err != nil {
		return nil, err
	}
	if err := q.packData(); err != nil {
		return nil, err
	}
	return q.commentList, nil
}

func (q *QueryCommentListFlow) checkData() error {
	if !model.NewUserInfoDAO().IsUserExistById(q.userId) {
		return fmt.Errorf("用户%d处于登出状态", q.userId)
	}
	if !model.NewVideoDAO().IsVideoExistById(q.videoId) {
		return fmt.Errorf("视频%d不能存在或已经被删除", q.videoId)
	}
	return nil
}

func (q *QueryCommentListFlow) prepareData() error {
	err := model.NewCommentDAO().QueryCommentListByVideoId(q.videoId, &q.comments)
	if err != nil {
		return err
	}
	// 根据前端的要求填充正确的时间格式
	err = util.FillCommentListFields(&q.comments)
	if err != nil {
		return err
	}
	return nil
}

func (q *QueryCommentListFlow) packData() error {
	q.commentList = &List{Comments: q.comments}
	return nil
}
