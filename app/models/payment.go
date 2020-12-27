package models

import "github.com/jinzhu/gorm"

type Payment struct {
	gorm.Model
	TransactionID int64 `json:"transactionId"`
	isPaid        bool  `json:"isPaid"`
	userID        int64 `json:"userId"`
}