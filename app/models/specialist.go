package models

import "github.com/jinzhu/gorm"

type Specialist struct {
	gorm.Model
	User    User
	UserID  int64    `json:"userId"`
	Parties []Party  `json:"party" gorm:"many2many:party_specialists"`
}