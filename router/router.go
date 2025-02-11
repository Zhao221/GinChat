package router

import (
	"GINCHAT/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/user/getUserList", service.GetUserList)
	return r
}
