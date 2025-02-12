package service

import (
	"GINCHAT/models"
	"GINCHAT/utils"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"math/rand"
	"net/http"
	"time"
)

// GetUserList
// @Summary 获取用户列表
// @Tags 用户模块
// @Success 200 {string} json{"code","message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.User, 0)
	data = models.GetUserList()
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @param user body models.User true "用户信息"
// @Success 200 {string} json{"code","message"}
// @Router /user/createUser [post]
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
	}
	data1 := models.FindUserByName(user.Name)
	if data1.Name != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "用户已注册",
		})
	}
	data2 := models.FindUserByEmail(user.Email)
	if data2.Name != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "用户已注册",
		})
	}
	data3 := models.FindUserByPhone(user.Phone)
	if data3.Name != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "用户已注册",
		})
	}
	// 用户密码加密
	// 设置随机数种子，以确保每次运行程序时生成不同的随机数
	rand.Seed(time.Now().UnixNano())
	// 生成一个 0 到 999999 之间的随机整数
	randomNum := rand.Intn(1000000)
	salt := fmt.Sprintf("%06d", randomNum)
	user.Salt = salt
	user.Password = utils.MakePassword(user.Password, salt)
	models.CreateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
	})
}

// LoginUserByNameAnPwd
// @Summary 用户登录
// @Tags 用户模块
// @param user body models.User true "用户信息"
// @Success 200 {string} json{"code","message"}
// @Router /user/loginUserByNameAnPwd [post]
func LoginUserByNameAnPwd(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
	}
	data := models.FindUserByName(user.Name)
	if data.Name != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "用户不存在",
		})
	}
	isCorrect := utils.ValidPassword(user.Password, data.Salt, data.Password)
	if !isCorrect {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "密码输入错误",
		})
	}
	models.LoginUserByNameAnPwd(user.Name, user.Password)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
	})
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @param id path uint true "用户ID"
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteUser [delete]
func DeleteUser(c *gin.Context) {
	var userID uint
	if err := c.ShouldBind(&userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
	}
	models.DeleteUser(userID)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
	})
}

// UpdateUser
// @Summary 更新用户
// @Tags 用户模块
// @param user body models.User true "用户信息"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [put]
func UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
	}
	// 手机号邮箱校验
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
	}
	models.UpdateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
	})
}

// 防止跨域站点伪造请求
var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMessage(c *gin.Context) {
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)
	MsgHandler(ws, c)
	fmt.Println(err)
}

func MsgHandler(ws *websocket.Conn, c *gin.Context) {
	for {
		msg, err := utils.Subscribe(c, utils.PublishKey)
		if err != nil {
			fmt.Println(err)
		}
		tm := time.Now().Format("2006-01-02 15:04:05")
		m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
		err = ws.WriteMessage(websocket.TextMessage, []byte(m))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func SendUserMessage(c *gin.Context) {
	models.Chat(c.Writer, c.Request)
}
