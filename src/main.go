package main

import (
	"backend-challenge/src/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80" // Porta padrão
	}

	http.HandleFunc("/", handlers.HandleContactForm)
	log.Printf("Starting server on port :%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
