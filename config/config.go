package config

import (
	"os"

	"github.com/joho/godotenv"
)

type appSecrets struct {
	Port string
	DbUrl string
}

var AppEnvs *appSecrets


func LoadEnv(){
	err := godotenv.Load()

	if err != nil {
		panic("Error loading env file")
	}

	AppEnvs = &appSecrets{
		Port: os.Getenv("PORT"),
	}
}