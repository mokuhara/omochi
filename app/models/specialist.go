package models

import "github.com/jinzhu/gorm"

type Specialist struct {
	gorm.Model
	User User
	UserID     int64  `json:"userId"`
}