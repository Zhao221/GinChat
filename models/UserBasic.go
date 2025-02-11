package models

import (
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string `gorm:"type:varchar(20);not null" json:"name"`
	Password      string `gorm:"type:varchar(255);not null" json:"password"`
	Phone         string `gorm:"type:varchar(20);not null" json:"phone"`
	Email         string `gorm:"type:varchar(20);not null" json:"email"`
	Identity      string `gorm:"type:varchar(20);not null" json:"identity"`
	ClientIP      string `gorm:"type:varchar(20);not null" json:"client_ip"`
	ClientPort    string `gorm:"type:varchar(20);not null" json:"client_port"`
	DeviceInfo    string `gorm:"type:varchar(20);not null" json:"device_info"`
	LoginTime     uint64 `gorm:"type:varchar(20);not null" json:"login_time"`
	HeartBeatTime uint64 `gorm:"type:varchar(20);not null" json:"heart_beat_time"`
	LogoutTime    uint64 `gorm:"type:varchar(20);not null" json:"logout_time"`
	IsLogout      bool   `gorm:"type:varchar(20);not null" json:"is_logout"`
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
