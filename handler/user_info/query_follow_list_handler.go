package user_info

import (
	"dousheng/model"
	"dousheng/service/user_info"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FollowListResponse struct {
	model.CommonResponse
	*user_info.FollowList
}

func QueryFollowListHandler(c *gin.Context) {
	NewProxyQueryFollowList(c).Do()
}

type ProxyQueryFollowList struct {
	*gin.Context

	userId int64

	*user_info.FollowList
}

func NewProxyQueryFollowList(context *gin.Context) *ProxyQueryFollowList {
	return &ProxyQueryFollowList{Context: context}
}

func (p *ProxyQueryFollowList) Do() {
	var err error
	if err = p.parseData(); err != nil {
		p.SendError(err.Error())
		return
	}
	if err = p.prepareData(); err != nil {
		p.SendError(err.Error())
		return
	}
	p.SendOk("请求成功")
}

func (p *ProxyQueryFollowList) parseData() error {
	rawUserId, _ := p.Get("user_id")
	userId, ok := rawUserId.(int64)
	if !ok {
		return errors.New("userId解析出错")
	}
	p.userId = userId
	return nil
}

func (p *ProxyQueryFollowList) prepareData() error {
	list, err := user_info.QueryFollowList(p.userId)
	if err != nil {
		return err
	}
	p.FollowList = list
	return nil
}

func (p *ProxyQueryFollowList) SendError(msg string) {
	p.JSON(http.StatusOK, FollowListResponse{
		CommonResponse: model.CommonResponse{
			StatusCode: 1,
			StatusMsg:  msg,
		},
	})
}

func (p *ProxyQueryFollowList) SendOk(msg string) {
	p.JSON(http.StatusOK, FollowListResponse{
		CommonResponse: model.CommonResponse{
			StatusCode: 0,
			StatusMsg:  msg,
		},
		FollowList: p.FollowList,
	})

}