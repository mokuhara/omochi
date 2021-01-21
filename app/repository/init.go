package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	_ "omochi/app/models"
	"omochi/config"
)

var DB = DBCon()

func DBCon() *gorm.DB{
	db, err := gorm.Open("mysql", config.Config.DbDSN)

	if err != nil {
		log.Fatalln(err)
	}

	//db.AutoMigrate(&models.User{}, &models.Client{}, &models.Specialist{}, &models.Company{})
	//db.AutoMigrate(&models.UserInfo{})
	//db.AutoMigrate(&models.Bizpack{}, &models.Party{}, &models.Product{}, &models.Category{})
	//db.AutoMigrate(&models.Payment{}, &models.Transaction{}, &models.VideoMeeting{}, models.Review{})
	//db.AutoMigrate(&models.Issue{})
	return db
}

func DBClose() {
	if err := DB.Close(); err != nil {
		panic(err)
	}
}
