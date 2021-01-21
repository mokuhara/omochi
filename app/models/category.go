package models

import _ "github.com/jinzhu/gorm"

type categoryType int

const (
	selection categoryType = iota + 1
	implement
	operation
)

type Category struct {
	Type categoryType `json:"type"`
}