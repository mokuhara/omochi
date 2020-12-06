package models

import "github.com/jinzhu/gorm"

type Party struct {
	gorm.Model
	Specialists []Specialist `json:"users" gorm:"many2many:party_specialists"`
	Bizpack []Bizpack `json:"bizpacks"`
}