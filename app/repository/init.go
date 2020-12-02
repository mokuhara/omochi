package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"omochi/app/models"
	"omochi/config"
)

func DBCon() *gorm.DB {
	db, err := gorm.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.User{}, &models.Client{}, &models.Specialist{})
	db.AutoMigrate(&models.UserInfo{})
	return db
}
