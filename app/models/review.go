package models

import "github.com/jinzhu/gorm"

type Review struct {
	gorm.Model
	TransactionID int64   `json:"transactionId"`
	UserID        int64   `gorm:"unique" json:"userId"`
	Message       string  `json:"message"`
	Rating        int64   `json:"rating"`
}