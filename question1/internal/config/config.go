package config

import (
	"github.com/joho/gotdotenv"
	"log"
)
const (
	WINDOW_SIZE = 10
)

func LoadEnv(){
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

}