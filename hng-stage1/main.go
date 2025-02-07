package main

import (
	"github.com/cozyCodr/hng-stage0/all_handlers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()

	// CORS configuration
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET"}),
		handlers.AllowedOrigins([]string{"*"}),
	)

	router.HandleFunc("/api/classify-number", all_handlers.ClassifyNumberHandler).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, cors(router)); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
