package video

import (
	"dousheng/middleware"
	"dousheng/model"
	"dousheng/service/video"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	model.CommonResponse
	*video.FeedVideoList
}

type ProxyFeedVideoList struct {
	*gin.Context
}

func NewProxyFeedVideoList(c *gin.Context) *ProxyFeedVideoList {
	return &ProxyFeedVideoList{Context: c}
}

func FeedVideoListHandler(c *gin.Context) {
	p := NewProxyFeedVideoList(c)
	token, ok := c.GetQuery("token")
	if !ok {
		err := p.DoNoToken()
		if err != nil {
			p.FeedVideoListError(err.Error())
		}
		return
	}

	// 有登陆状态
	err := p.DoHasToken(token)
	if err != nil {
		p.FeedVideoListError(err.Error())
	}
}

// DoNoToken 未登录的视频流推送处理
func (p *ProxyFeedVideoList) DoNoToken() error {
	rawTimestamp := p.Query("latest_time")
	var latestTime time.Time
	intTime, err := strconv.ParseInt(rawTimestamp, 10, 64)
	if err == nil {
		latestTime = time.Unix(0, intTime*1e6) // 注意：前端传来的时间戳是以ms为单位的
	}
	videoList, err := video.QueryFeedVideoList(0, latestTime)
	if err != nil {
		return err
	}
	p.FeedVideoListOk(videoList)
	return nil
}

// DoHasToken 如果是登陆状态，则生成UserId字段
func (p *ProxyFeedVideoList) DoHasToken(token string) error {
	// 解析成功
	if claim, ok := middleware.ParseToken(token); ok {
		// token超时
		if time.Now().Unix() > claim.ExpiresAt {
			return errors.New("token超时")
		}
		rawTimestamp := p.Query("latest_time")
		var latestTime time.Time
		intTime, err := strconv.ParseInt(rawTimestamp, 10, 64)
		if err == nil {
			latestTime = time.Unix(0, intTime*1e6) // 注意：前端传来的时间戳是以ms为单位的
		}
		// 调用service层接口
		videoList, err := video.QueryFeedVideoList(claim.UserId, latestTime)
		if err != nil {
			return err
		}
		p.FeedVideoListOk(videoList)
		return nil
	}
	// 解析失败
	return errors.New("token不正确")
}

func (p *ProxyFeedVideoList) FeedVideoListError(msg string) {
	p.JSON(http.StatusOK, FeedResponse{CommonResponse: model.CommonResponse{
		StatusCode: 1,
		StatusMsg:  msg,
	}})
}

func (p *ProxyFeedVideoList) FeedVideoListOk(videoList *video.FeedVideoList) {
	p.JSON(http.StatusOK, FeedResponse{
		CommonResponse: model.CommonResponse{
			StatusCode: 0,
		},
		FeedVideoList: videoList,
	})
}
