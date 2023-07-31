package config

import (
	"github.com/joho/godotenv"
	"log"
)

//LoadEnvVariables load .Env file
func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
