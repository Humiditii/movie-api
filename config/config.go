package config

import (
	"os"

	"github.com/joho/godotenv"
)

type appSecrets struct {
	Port string
	DbUrl string
	DatabaseName string
}

var AppEnvs *appSecrets


func LoadEnv(){
	err := godotenv.Load()

	if err != nil {
		panic("Error loading env file")
	}

	AppEnvs = &appSecrets{
		Port: os.Getenv("PORT"),
		DbUrl: os.Getenv("DB_URL"),
		DatabaseName: os.Getenv("DATABASE_NAME"),
	}
}