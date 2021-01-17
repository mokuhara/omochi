package models

import (
	"github.com/jinzhu/gorm"
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
	StartAt string `json:"startAt"`
	EndAt string `json:"endAt"`
	ApplicationDeadline string `json:"applicationDeadline"`
	UserID int64 `json:"userId"`
}