package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type VideoMeeting struct {
	gorm.Model
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	StartedAt time.Time `json:"startedAt"`
}