package router

import (
	"GINCHAT/docs"
	"GINCHAT/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/user/getUserList", service.GetUserList)
	r.POST("/user/createUser", service.CreateUser)
	r.DELETE("/user/deleteUser", service.DeleteUser)
	r.PUT("/user/updateUser", service.UpdateUser)
	r.POST("/user/loginUserByNameAnPwd", service.LoginUserByNameAnPwd)
	r.POST("/user/senMessage", service.SendMessage)
	return r
}
