package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-challenge/http/server/handler"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(handler.LogginMiddleware)
	router.HandleFunc("/note/{noteId}", handler.HTTPGetNote).Methods("GET")

	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 3 * time.Second,
	}
	fmt.Println("Starting http server...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server startup failure", err)
	}
}
