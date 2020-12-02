package models

import "github.com/jinzhu/gorm"

type Type int

const (
	specialist Type = iota + 1
	client
)

type User struct {
	gorm.Model
	Email      string `gorm:"unique" json:"email"`
	Password   string `json:"password"`
	Type       Type  `json:"type"`
}