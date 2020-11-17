package config

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
)

type ConfigList struct  {
	DbName string
	SQLDriver string

	Port int
	LogFile string
}

var Config ConfigList

func init(){
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: #{err}")
		os.Exit(1)
	}

	Config = ConfigList{
		DbName: cfg.Section("db").Key("name").String(),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		Port: cfg.Section("web").Key("port").MustInt(),
		LogFile: cfg.Section("web").Key("log_file").String(),
	}
}