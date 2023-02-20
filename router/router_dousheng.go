package router

import (
	"dousheng/handler/user_info"
	"dousheng/handler/user_login"
	"dousheng/handler/video"
	"dousheng/middleware"
	"dousheng/model"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	model.InitDB()
	r := gin.Default()

	r.Static("static", "./static")

	baseGroup := r.Group("/douyin")
	// 根据灵活性考虑是否加入JWT中间件来进行鉴权，还是在之后再做鉴权
	baseGroup.GET("/feed/", video.FeedVideoListHandler)
	baseGroup.GET("/user/", middleware.JWTMiddleWare(), user_info.UserInfoHandler)
	baseGroup.POST("/user/login/", middleware.SHAMiddleWare(), user_login.UserLoginHandler)
	baseGroup.POST("/user/register/", middleware.SHAMiddleWare(), user_login.UserRegisterHandler)
	baseGroup.POST("/publish/action/", middleware.JWTMiddleWare(), video.PublishVideoHandler)
	baseGroup.GET("/publish/list/", middleware.JWTMiddleWare(), video.QueryVideoListHandler)

	return r
}
