package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"omochi/config"
)

func DBCon() *gorm.DB {
	db, err := gorm.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}
	//db.AutoMigrate(&models.User{}, &models.Client{}, &models.Specialist{}, &models.Company{})
	//db.AutoMigrate(&models.UserInfo{})
	//db.AutoMigrate(&models.Bizpack{}, &models.Party{}, &models.Product{}, &models.Category{})
	return db
}
