package main

import (
	"log"
	"net/http"

	"github.com/lordofthemind/course_webSockets_Go/internal/handlers"
)

func main() {
	mux := routes()

	log.Println("Starting channel listener")
	go handlers.ListenToWsChannel()

	log.Println("Starting server on :9090")

	_ = http.ListenAndServe(":9090", mux)
}
