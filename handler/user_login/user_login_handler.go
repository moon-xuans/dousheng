package user_login

import (
	"dousheng/model"
	"dousheng/service/user_login"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserLoginResponse struct {
	model.CommonResponse
	*user_login.LoginResponse
}

func UserLoginHandler(c *gin.Context) {
	username := c.Query("username")
	raw, _ := c.Get("password")
	password, ok := raw.(string)
	if !ok {
		c.JSON(http.StatusOK, UserLoginResponse{
			CommonResponse: model.CommonResponse{
				StatusCode: 1,
				StatusMsg:  "密码解析错误",
			},
		})
		return
	}
	userLoginResponse, err := user_login.QueryUserLogin(username, password)

	// 用户不存在返回对应的错误
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			CommonResponse: model.CommonResponse{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	// 用户存在，返回对应的id和token
	c.JSON(http.StatusOK, UserLoginResponse{
		CommonResponse: model.CommonResponse{StatusCode: 0},
		LoginResponse:  userLoginResponse,
	})
}
