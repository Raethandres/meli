package handler

import (
	"net/http"

	. "github.com/patrickmn/go-cache"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func Routes(c *Cache) *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/", NothingToDoHandler).Methods("GET")
	r.HandleFunc("/topsecret/", TopSecretHandler).Methods("POST")
	r.HandleFunc("/topsecret_split/{name}", func(w http.ResponseWriter, r *http.Request) {
		TopSecretSplitGetHandler(w, r, c)
	}).Methods("GET")
	r.HandleFunc("/topsecret_split/", func(w http.ResponseWriter, r *http.Request) {
		TopSecretSplitGetHandler(w, r, c)
	}).Methods("POST")

	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "../../swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	r.Handle("/docs", sh)
	r.Handle("/swagger.yaml", http.FileServer(http.Dir("../../")))

	return r
}
