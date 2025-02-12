package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	OwnerId     uint   `json:"owner_id"`  // 谁的关系信息
	TargetId    uint   `json:"target_id"` // 对应的谁
	Type        int    `json:"type"`      // 对应的类型 0 1 3
	Description string `son:"description"`
}

func (c *Contact) TableName() string {
	return "contact"
}
