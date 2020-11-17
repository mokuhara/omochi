package main

import (
	"omochi/config"
	"omochi/utils"
)

func main(){
	utils.LoggingSettings(config.Config.LogFile)
}