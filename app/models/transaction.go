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
	ClientUserID int64 `json:"clientUserId"`
	SpecialistUserID int64 `json:"specialistUserId"`
	BizpackID int64`json:"bizpackId" gorm:"default:null"`
	Bizpack Bizpack `json:"bizpack" gorm:"save_associations:false"`
	IssueID int64 `json:"issueId" gorm:"default:null"`
	Issue Issue `json:"issue" gorm:"save_associations:false"`
	Status        status         `json:"status"`
	Title string `json:"title"`
	Category    Category  `json:"category"`
	Description string `json:"description"`
	UnitPrice   int64     `json:"unitPrice"`
	Duration    int64     `json:"duration"`
	ClientAcceptance int64 `json:"clientAcceptance"`
	SpecialistAcceptance int64 `json:"specialistAcceptance"`
	VideoMeetings []VideoMeeting `json:"videoMeetings" gorm:"save_associations:false"`
	//Payments      []Payment      `json:"payments" gorm:"save_associations:false"`
	Reviews       []Review       `json:"reviews" gorm:"save_associations:false"`
}