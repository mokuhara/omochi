package config

import (
	"github.com/joho/godotenv"
	"fmt"
	"log"
	"os"
)

type ConfigList struct {
	DbName string
	DbDSN string

	Port string
	Origins []string
	LogFile string

	ZoomEndpoint string
	ZoomJwt string
}

var Config ConfigList

func init() {

	if os.Getenv("GO_ENV") == "production" {
		// load .env file
		err := godotenv.Load()

		if err != nil {
			log.Printf("Failed to loading .env file")
			os.Exit(1)
		}
	}

	// gormでDBに接続するための情報
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	Config = ConfigList{
		DbDSN: dsn,
		Port: os.Getenv("PORT"),
		LogFile: os.Getenv("LOG_FILE"),
		Origins: []string{ os.Getenv("WEB_ORIGIN") },
		ZoomEndpoint: os.Getenv("ZOOM_ENDPOINT"),
		ZoomJwt: os.Getenv("ZOOM_JWT"),
	}
}
