package handler

import (
	"encoding/json"
	. "meli/pkg/request"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/patrickmn/go-cache"
)

type Empty struct{}

func TopSecretSplitPushSingleHandler(w http.ResponseWriter, r *http.Request, c *Cache) {
	var request SateliteRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		BadRequestMessage(w, "Empty Body")
		return
	}
	vars := mux.Vars(r)
	request.Name = vars["satellite_name"]

	c.Set(request.Name, request, DefaultExpiration)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Empty{})
}
