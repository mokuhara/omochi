package main

import (
	"omochi/app/repository"
	"omochi/config"
	"omochi/routes"
	"omochi/utils"
)

func main(){
	utils.LoggingSettings(config.Config.LogFile)
	routes.Router()
	repository.DBClose()
}