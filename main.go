package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Message 	string 		`json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Version 	string 		`json:"version"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := Response{
		Message: "OK",
		Timestamp: time.Now().UTC(),
		Version: "1.0.0",
	}

	json.NewEncoder(w).Encode(response)
}

func rootHandler(w http.ResponseWriter,  r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := Response{
		Message: "Hello from Go API",
		Timestamp: time.Now().UTC(),
		Version: "1.0.0",
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/", rootHandler)

	log.Printf("Server starting on port %s...", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
