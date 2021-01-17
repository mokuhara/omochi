package main

import (
	"omochi/app/repository"
	"omochi/routes"
	// "omochi/config"
	// "omochi/utils"
)

func main(){
	// utils.LoggingSettings(config.Config.LogFile)
	routes.Router()
	repository.DBClose()
}