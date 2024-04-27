package main

import (
	"log"
	"net/http"

	Handler "scale-x/Handler"
	Service "scale-x/Service"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Create a new mux router
	mux := mux.NewRouter()

	// Initialize the service
	service := Service.NewService()

	// Create handlers and pass the service
	handlers := Handler.Handlers{Service: *service}

	// Define endpoints and corresponding handler functions

	// Login endpoint
	mux.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	// Home endpoint
	mux.HandleFunc("/home", handlers.HomeHandler).Methods("GET")

	// Add book endpoint
	mux.HandleFunc("/addBook", handlers.AddBookHandler).Methods("POST")

	// Delete book endpoint
	mux.HandleFunc("/deleteBook", handlers.DeleteBookHandler).Methods("DELETE")

	// Create a CORS middleware with allowed origins, methods, and headers
	router := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Add PUT here
		AllowedHeaders: []string{"*"},
	}).Handler(mux)

	// Specify the address to listen on
	listenAddr := ":8080"

	// Log the address to the console
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s", listenAddr, listenAddr)

	// Start the HTTP server with the CORS middleware and router
	log.Fatal(http.ListenAndServe(listenAddr, router))
}
