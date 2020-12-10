package models

import "github.com/jinzhu/gorm"

type categoryType int

const (
	selection categoryType = iota + 1
	implement
	operation
)

type Category struct {
	gorm.Model
	Type categoryType `json:"type"`
	BizpackID int64
}