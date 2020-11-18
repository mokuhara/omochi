package main

import (
	"omochi/app/models"
	"omochi/config"
	"omochi/utils"
)

func main(){
	utils.LoggingSettings(config.Config.LogFile)
	models.DBCon()
}