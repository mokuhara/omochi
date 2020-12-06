package models

import "github.com/jinzhu/gorm"

type Bizpack struct {
	gorm.Model
	Products    []Product `json:"products"`
	User        User
	UserID      int64     `json:"userId"`
	Industry    string    `json:"industry"`
	Scale       int64     `json:"scale"`
	Description string    `json:"description"`
	UnitPrice   int64     `json:"unitPrice"`
	Duration    int64     `json:"duration"`
	IsPublic    bool      `json:"isPublic"`
	Party       Party
	PartyID     int64	  `json:"partyId"`
}