package handler

import (
	"log"
	"os"

	"encoding/json"
	. "meli/pkg/response"
	"net/http"

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

func BadRequest(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	response := ErrorResponse{Error: msg}
	json.NewEncoder(w).Encode(response)
}
