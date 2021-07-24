package handler

import (
	"encoding/json"
	. "meli/pkg/response"
	"net/http"
)

func BadRequestMessage(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	response := ErrorResponse{Error: msg}
	json.NewEncoder(w).Encode(response)
}
