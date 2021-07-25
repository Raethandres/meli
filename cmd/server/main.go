package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"meli/pkg/handler"

	"github.com/patrickmn/go-cache"
)

func main() {
	handler.GoDotEnv()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	c := cache.New(5*time.Minute, 10*time.Minute)
	infoLog.Printf("Starting server on %s", os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), handler.Routes(c))
}
