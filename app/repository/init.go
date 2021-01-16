package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	_ "omochi/app/models"
	"omochi/config"
)

var DB = DBCon()

func DBCon() *gorm.DB{
	DB, err := gorm.Open(config.Config.SQLDriver, config.Config.DbName)

	if err != nil {
		log.Fatalln(err)
	}

	//DB.AutoMigrate(&models.User{}, &models.Client{}, &models.Specialist{}, &models.Company{})
	//DB.AutoMigrate(&models.UserInfo{})
	//DB.AutoMigrate(&models.Bizpack{}, &models.Party{}, &models.Product{}, &models.Category{})
	//DB.AutoMigrate(&models.Payment{}, &models.Transaction{}, &models.VideoMeeting{}, models.Review{})
	//DB.AutoMigrate(&models.Issue{})
	return DB
}

func DBClose() {
	if err := DB.Close(); err != nil {
		panic(err)
	}
}
