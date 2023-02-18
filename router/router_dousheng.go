package router

import (
	"dousheng/handler/user_login"
	"dousheng/handler/video"
	"dousheng/model"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	model.InitDB()
	r := gin.Default()

	baseGroup := r.Group("/douyin")

	baseGroup.POST("/user/login/", user_login.UserLoginHandler)
	baseGroup.POST("/user/register/", user_login.UserRegisterHandler)
	baseGroup.POST("/publish/action", video.PublishVideoHandler)

	return r
}
