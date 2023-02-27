package video

import (
	"dousheng/model"
	"dousheng/service/video"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostFavorHandler(c *gin.Context) {
	NewProxyPostFavorHandler(c).Do()
}

type ProxyFavorHandler struct {
	*gin.Context

	userId     int64
	videoId    int64
	actionType int64
}

func NewProxyPostFavorHandler(c *gin.Context) *ProxyFavorHandler {
	return &ProxyFavorHandler{Context: c}
}

func (p *ProxyFavorHandler) Do() {
	// 解析参数
	if err := p.parseData(); err != nil {
		p.SendError(err.Error())
		return
	}
	// 正式调用
	err := video.PostFavorState(p.userId, p.videoId, p.actionType)
	if err != nil {
		p.SendError(err.Error())
		return
	}
	// 成功返回
	p.SendOk()
}

func (p *ProxyFavorHandler) parseData() error {
	// 解析userId
	rawUserId, _ := p.Get("user_id")
	userId, ok := rawUserId.(int64)
	if !ok {
		return errors.New("userId解析出错")
	}

	rawVideoId := p.Query("video_id")
	videoId, err := strconv.ParseInt(rawVideoId, 10, 64)
	if err != nil {
		return err
	}
	rawActionType := p.Query("action_type")
	actionType, err := strconv.ParseInt(rawActionType, 10, 64)
	if err != nil {
		return err
	}
	p.videoId = videoId
	p.actionType = actionType
	p.userId = userId
	return nil
}

func (p *ProxyFavorHandler) SendError(msg string) {
	p.JSON(http.StatusOK, model.CommonResponse{
		StatusCode: 1,
		StatusMsg:  msg,
	})
}

func (p *ProxyFavorHandler) SendOk() {
	p.JSON(http.StatusOK, model.CommonResponse{
		StatusCode: 0,
	})

}
