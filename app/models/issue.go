package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Issue struct {
	gorm.Model
	Title string `json:"title"`
	Background string `json:"background"`
	Description string `json:"description"`
	DesiredSpecialist string `json:"desiredSpecialist"`
	RequiredItem string `json:"requiredItem"`
	ClientInfo string `json:"clientInfo"`
	Category Category
	CategoryID int64 `json:"categoryId"`
	Budget int64 `json:"budget"`
	RecruitmentCapacity int64 `json:"recruitmentCapacity"`
	StartAt time.Time `json:"startAt"`
	EndAt time.Time `json:"endAt"`
	ApplicationDeadline time.Time `json:"applicationDeadline"`
	UserID int64 `json:"userId"`
}