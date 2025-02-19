package router

import (
	"Chat/docs"
	"Chat/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	// swag
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 静态资源
	r.Static("/asset", "asset/")
	r.LoadHTMLGlob("views/**/*")

	// 首页
	r.GET("/", service.GetIndex)
	r.GET("/index", service.GetIndex)
	r.GET("/toRegister", service.ToRegister)

	// 用户模块
	r.GET("/user/getUserList", service.GetUserList)
	r.POST("/user/createUser", service.CreateUser)
	r.DELETE("/user/deleteUser", service.DeleteUser)
	r.PUT("/user/updateUser", service.UpdateUser)
	r.POST("/user/findUserByNameAndPwd", service.LoginUserByNameAnPwd)
	r.POST("/user/senMessage", service.SendMessage)
	r.POST("/user/sendUserMessage", service.SendUserMessage)
	return r
}
