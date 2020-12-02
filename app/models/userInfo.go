package models

import "github.com/jinzhu/gorm"

type motivation int

const (
	soon motivation = iota + 1
	consideration
	justLooking
)

type UserInfo struct {
	gorm.Model
	User User
	UserID int64 `json:"userId"`
	Name string `json:"name"`
	kana string `json:"kana"`
	Phone string `json:"phone"`
	CompanyName string `json:"companyName"`
	Department string `json:"department"`
	Position string `json:"position"`
	CompanyPhone string `json:"companyPhone"`
	Motivation motivation `json:"motivation"`
	SupportRequest bool `json:"supportRequest"`
	Consent bool `json:"consent"`
}