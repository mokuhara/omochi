package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name string `json:"name"`
	BizpackID int64
	// PortfolioID int64
}
