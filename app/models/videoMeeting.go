package models

import (
	"github.com/jinzhu/gorm"
)

type VideoMeeting struct {
	gorm.Model
	Name      string    `json:"topic"`
	Url       string    `json:"join_url"`
	StartedAt string `json:"start_time"`
	TransactionID int64 `json:"transactionId"`
}