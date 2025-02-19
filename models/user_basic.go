package models

import (
	"GINCHAT/utils"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Salt          string    `bind:"-" gorm:"type:varchar(255);not null" json:"salt"` // 加密随机数
	Name          string    `bind:"-" gorm:"type:varchar(255);not null" json:"name"`
	Password      string    `bind:"-" gorm:"type:varchar(255);not null" json:"password"`
	Phone         string    `bind:"-" gorm:"type:varchar(255);not null" json:"phone" valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string    `bind:"-" gorm:"type:varchar(255);not null" json:"email" valid:"email"`
	Identity      string    `bind:"-" gorm:"type:varchar(255);not null" json:"identity"`
	ClientIP      string    `bind:"-" gorm:"type:varchar(255);not null" json:"client_ip"`
	ClientPort    string    `bind:"-" gorm:"type:varchar(255);not null" json:"client_port"`
	DeviceInfo    string    `bind:"-" gorm:"type:varchar(255);not null" json:"device_info"`
	LoginTime     time.Time `gorm:"type:datetime;default:null" json:"login_time"`
	HeartBeatTime time.Time `gorm:"type:datetime;default:null" json:"heart_beat_time"`
	LoginOutTime  time.Time `gorm:"type:datetime;default:null" json:"login_out_time"`
	IsLogout      bool      `bind:"-" gorm:"type:bool" json:"is_logout"`
}

func (u *User) TableName() string {
	return "user"
}

func GetUserList() []*User {
	userList := make([]*User, 0)
	utils.DB.Find(&userList)
	return userList
}

func FindUserByName(name string) User {
	var user User
	utils.DB.Where("name = ?", name).First(&user)
	return user
}

func FindUserByPhone(phone string) User {
	var user User
	utils.DB.Where("phone = ?", phone).First(&user)
	return user
}

func FindUserByEmail(email string) User {
	var user User
	utils.DB.Where("email = ?", email).First(&user)
	return user
}

func LoginUserByNameAnPwd(name string, password string) User {
	var user User
	utils.DB.Where("name = ? and password = ?", name, password).First(&user)
	// token加密
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := utils.MD5Encode(str)
	utils.DB.Model(&user).Where("id = ?", user.ID).Update("identity", temp)
	return user
}

func CreateUser(user User) *gorm.DB {
	return utils.DB.Create(&user)
}

func DeleteUser(userID uint) *gorm.DB {
	var user User
	return utils.DB.Where("id = ?", userID).Delete(&user)
}

func UpdateUser(user User) *gorm.DB {
	return utils.DB.Where("id = ?", user.ID).Updates(&user)
}
