package models

import "github.com/jinzhu/gorm"


type status int

const (
	preInterview status = iota + 1
	preAccept
	lostOrder
	preSendQuotation
	preSendApplicationForm
	contractComplete
	endContract
)


type Transaction struct {
	gorm.Model
	BizpackID     int64          `json:"bizpackId"`
	Status        status         `json:"status"`
	VideoMeetings []VideoMeeting `json:"videoMeetings" gorm:"save_associations:false"`
	Payments      []Payment      `json:"payments" gorm:"save_associations:false"`
	Reviews       []Review       `json:"reviews" gorm:"save_associations:false"`
}