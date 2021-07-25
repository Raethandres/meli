package handler

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GoDotEnv() {

	// load .env file
	if os.Getenv("isDocker") != "true" {
		log.Print(os.Getenv("isDocker"), "os")
		err := godotenv.Load("../../.env")

		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

}
