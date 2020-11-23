package models

import "github.com/jinzhu/gorm"

type Profile struct {
	gorm.Model
	User        User
	UserID      int64  `json:"userId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}