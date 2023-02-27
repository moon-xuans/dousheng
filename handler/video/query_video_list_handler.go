package video

import (
	"dousheng/model"
	"dousheng/service/video"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VideosResponse struct {
	model.CommonResponse
	*video.Videos
}

func QueryVideoListHandler(c *gin.Context) {
	p := NewProxyQueryVideoList(c)
	rawId, _ := c.Get("user_id")
	err := p.DoQueryListByUserId(rawId)
	if err != nil {
		p.QueryVideoListError(err.Error())
	}
}

// ProxyQueryVideoList 代理类
type ProxyQueryVideoList struct {
	c *gin.Context
}

func NewProxyQueryVideoList(c *gin.Context) *ProxyQueryVideoList {
	return &ProxyQueryVideoList{c: c}
}

// DoQueryListByUserId 根据userId字段进行查询
func (p *ProxyQueryVideoList) DoQueryListByUserId(rawId interface{}) error {
	userId, ok := rawId.(int64)
	if !ok {
		return errors.New("userId解析出错")
	}

	videoList, err := video.QueryVideoListByUserId(userId)
	if err != nil {
		return err
	}

	p.QueryVideoListOk(videoList)
	return nil
}

func (p *ProxyQueryVideoList) QueryVideoListError(msg string) {
	p.c.JSON(http.StatusOK, VideosResponse{
		CommonResponse: model.CommonResponse{
			StatusCode: 1,
			StatusMsg:  msg,
		},
	})
}

func (p *ProxyQueryVideoList) QueryVideoListOk(videoList *video.Videos) {
	p.c.JSON(http.StatusOK, VideosResponse{
		CommonResponse: model.CommonResponse{
			StatusCode: 0,
		},
		Videos: videoList,
	})
}
