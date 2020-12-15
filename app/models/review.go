package models

import "github.com/jinzhu/gorm"

type Review struct {
	gorm.Model
	TransactionID int64   `json:"transactionId"`
	UserID        int64   `json:"userId"`
	Message       string  `json:"message"`
}