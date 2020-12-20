package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type VideoMeeting struct {
	gorm.Model
	Name      string    `json:"topic"`
	Url       string    `json:"join_url"`
	StartedAt time.Time `json:"start_time"`
}