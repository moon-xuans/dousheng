package router

import (
	"dousheng/model"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	model.InitDB()
	r := gin.Default()

	return r
}
