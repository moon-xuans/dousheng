package router

import (
	"dousheng/handler/comment"
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
	baseGroup.GET("/publish/list/", middleware.NoAuthToGetUserId(), video.QueryVideoListHandler)

	// 互动
	baseGroup.POST("/favorite/action/", middleware.JWTMiddleWare(), video.PostFavorHandler)
	baseGroup.GET("/favorite/list/", middleware.NoAuthToGetUserId(), video.QueryFavorVideoListHandler)
	baseGroup.POST("/comment/action/", middleware.JWTMiddleWare(), comment.PostCommentHandler)
	baseGroup.GET("/comment/list/", middleware.NoAuthToGetUserId(), comment.QueryCommentListHandler)

	// 社交
	baseGroup.POST("/relation/action/", middleware.JWTMiddleWare(), user_info.PostFollowActionHandler)
	baseGroup.GET("/relation/follow/list/", middleware.NoAuthToGetUserId(), user_info.QueryFollowListHandler)
	baseGroup.GET("/relation/follower/list/", middleware.NoAuthToGetUserId(), user_info.QueryFollowerListHandler)

	return r
}
