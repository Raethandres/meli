package handler

import (
	"encoding/json"
	"log"
	. "meli/pkg/request"
	"net/http"

	. "meli/cmd/data"

	"github.com/gorilla/mux"
	. "github.com/patrickmn/go-cache"
)

type Empty struct{}

func TopSecretSplitPushSingleHandler(w http.ResponseWriter, r *http.Request, c *Cache) {
	var request SateliteRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		BadRequest(w, "not found")
		return
	}
	vars := mux.Vars(r)
	request.Name = vars["name"]
	log.Print(request)
	c.Set(request.Name, request, DefaultExpiration)
	kenobi, kenobiExist := c.Get(request.Name)
	log.Print(kenobi, kenobiExist, KENOBI)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Empty{})
}
