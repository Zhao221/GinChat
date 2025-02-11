package service

import (
	"GINCHAT/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 0)
	data = models.GetUserList()
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
