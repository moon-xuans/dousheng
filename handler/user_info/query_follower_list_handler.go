package user_info

import (
	"dousheng/model"
	"dousheng/service/user_info"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FollowerListResponse struct {
	model.CommonResponse
	*user_info.FollowerList
}

func QueryFollowerListHandler(c *gin.Context) {
	NewProxyQueryFollowerList(c).Do()
}

type ProxyQueryFollowerList struct {
	*gin.Context

	userId int64

	*user_info.FollowerList
}

func NewProxyQueryFollowerList(context *gin.Context) *ProxyQueryFollowerList {
	return &ProxyQueryFollowerList{Context: context}
}

func (p *ProxyQueryFollowerList) Do() {
	var err error
	if err = p.parseData(); err != nil {
		p.SendError(err.Error())
		return
	}
	if err = p.prepareData(); err != nil {
		if errors.Is(err, user_info.ErrUserNotExist) {
			p.SendError(err.Error())
		} else {
			p.SendError(err.Error())
		}
		return
	}
	p.SendOk("请求成功")
}

func (p *ProxyQueryFollowerList) parseData() error {
	rawUserId, _ := p.Get("user_id")
	userId, ok := rawUserId.(int64)
	if !ok {
		return errors.New("userId解析出错")
	}
	p.userId = userId
	return nil
}

func (p *ProxyQueryFollowerList) prepareData() error {
	list, err := user_info.QueryFollowerList(p.userId)
	if err != nil {
		return err
	}
	p.FollowerList = list
	return nil
}

func (p *ProxyQueryFollowerList) SendError(msg string) {
	p.JSON(http.StatusOK, FollowListResponse{
		CommonResponse: model.CommonResponse{
			StatusCode: 1,
			StatusMsg:  msg,
		},
	})
}

func (p *ProxyQueryFollowerList) SendOk(msg string) {
	p.JSON(http.StatusOK, FollowerListResponse{
		CommonResponse: model.CommonResponse{
			StatusCode: 0,
			StatusMsg:  msg,
		},
		FollowerList: p.FollowerList,
	})

}
