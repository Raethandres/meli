package handler

import "net/http"

func NothingToDoHandler(w http.ResponseWriter, r *http.Request) {
	BadRequestMessage(w, "Nothing to do")
}