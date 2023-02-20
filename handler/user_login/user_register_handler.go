package user_login

import (
	"dousheng/model"
	"dousheng/service/user_login"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRegisterResponse struct {
	model.CommonResponse
	*user_login.LoginResponse
}

func UserRegisterHandler(c *gin.Context) {
	username := c.Query("username")
	rawVal, _ := c.Get("password")
	password, ok := rawVal.(string)
	if !ok {
		c.JSON(http.StatusOK, UserRegisterResponse{
			CommonResponse: model.CommonResponse{
				StatusCode: 1,
				StatusMsg:  "密码解析错误",
			},
		})
		return
	}
	registerResponse, err := user_login.PostUserLogin(username, password)
	if err != nil {
		c.JSON(http.StatusOK, UserRegisterResponse{
			CommonResponse: model.CommonResponse{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	c.JSON(http.StatusOK, UserRegisterResponse{
		CommonResponse: model.CommonResponse{StatusCode: 0},
		LoginResponse:  registerResponse,
	})
}
