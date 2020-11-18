package main

import (
	"omochi/app/repository"
	"omochi/config"
	"omochi/utils"
)

func main(){
	utils.LoggingSettings(config.Config.LogFile)
	repository.DBCon()
}