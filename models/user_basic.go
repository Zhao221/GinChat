package models

import (
	"GINCHAT/utils"
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string    `gorm:"type:varchar(255);not null" json:"name"`
	Password      string    `gorm:"type:varchar(255);not null" json:"password"`
	Phone         string    `gorm:"type:varchar(255);not null" json:"phone"`
	Email         string    `gorm:"type:varchar(255);not null" json:"email"`
	Identity      string    `gorm:"type:varchar(255);not null" json:"identity"`
	ClientIP      string    `gorm:"type:varchar(255);not null" json:"client_ip"`
	ClientPort    string    `gorm:"type:varchar(255);not null" json:"client_port"`
	DeviceInfo    string    `gorm:"type:varchar(255);not null" json:"device_info"`
	LoginTime     time.Time `gorm:"type:datetime;default:null" json:"login_time"`
	HeartBeatTime time.Time `gorm:"type:datetime;default:null" json:"heart_beat_time"`
	LoginOutTime  time.Time `gorm:"type:datetime;default:null" json:"login_out_time"`
	IsLogout      bool      `gorm:"type:bool" json:"is_logout"`
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	userList := make([]*UserBasic, 0)
	utils.DB.Find(&userList)
	return userList
}
