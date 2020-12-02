package models

import "github.com/jinzhu/gorm"

type Client struct {
	gorm.Model
	User User
	UserID     int64  `json:"userId"`
	Company Company
	CompanyID  int64  `json:"companyId"`
}