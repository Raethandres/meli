package handler

import "net/http"

func NothingToDoHandler(w http.ResponseWriter, r *http.Request) {
	BadRequest(w, "Nothing to do")
}
