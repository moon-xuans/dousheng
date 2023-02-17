package router

import (
	"dousheng/models"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	models.InitDB()
	r := gin.Default()

	return r
}
