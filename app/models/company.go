package models

import "github.com/jinzhu/gorm"

type Company struct {
	gorm.Model
	Name string `json:"name"`
	Clients []Client `json:"client" gorm:"association_autocreate:false"`
}