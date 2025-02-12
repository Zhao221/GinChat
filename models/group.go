package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name        string `json:"name"`
	OwnerID     uint   `json:"owner_id"`
	Icon        string `json:"icon"`
	Type        int    `json:"type"`
	Description string `json:"description"`
}

func (g *Group) TableName() string {
	return "group"
}
