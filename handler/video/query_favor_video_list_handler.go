package video

import (
	"dousheng/model"
	"dousheng/service/video"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FavorVideoListResponse struct {
	model.CommonResponse
	*video.FavorList
}

func QueryFavorVideoListHandler(c *gin.Context) {
	NewProxyFavorVideoListHandler(c).Do()
}

type ProxyFavorVideoListHandler struct {
	*gin.Context

	userId int64
}

func NewProxyFavorVideoListHandler(c *gin.Context) *ProxyFavorVideoListHandler {
	return &ProxyFavorVideoListHandler{Context: c}
}

func (p *ProxyFavorVideoListHandler) Do() {
	// 解析参数
	if err := p.parseData(); err != nil {
		p.SendError(err.Error())
		return
	}

	// 正式调用
	favorVideoList, err := video.QueryFavorVideoList(p.userId)
	if err != nil {
		p.SendError(err.Error())
		return
	}
	// 成功返回
	p.SendOk(favorVideoList)
}

func (p *ProxyFavorVideoListHandler) parseData() error {
	rawUserId, _ := p.Get("user_id")
	userId, ok := rawUserId.(int64)
	if !ok {
		return errors.New("userId解析出错")
	}
	p.userId = userId
	return nil
}

func (p *ProxyFavorVideoListHandler) SendError(msg string) {
	p.JSON(http.StatusOK, FavorVideoListResponse{
		CommonResponse: model.CommonResponse{
			StatusCode: 1,
			StatusMsg:  msg,
		},
	})
}

func (p *ProxyFavorVideoListHandler) SendOk(favorList *video.FavorList) {
	p.JSON(http.StatusOK, FavorVideoListResponse{
		CommonResponse: model.CommonResponse{StatusCode: 0},
		FavorList:      favorList,
	})
}
