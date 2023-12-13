package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"test-signer/api"
	"test-signer/database"
	"test-signer/middleware"
	"test-signer/store"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve the JWT secret key
	jwtKey := os.Getenv("JWT_SECRET_KEY")
	if jwtKey == "" {
		log.Fatal("JWT_SECRET_KEY must be set in the .env file")
	}

	// Initialize the database
	db, err := database.InitializeDatabase("your_connection_string")
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	// Create a new store
	store := store.NewStore(db)

	// Set up the HTTP server
	http.HandleFunc("/sign", middleware.JWTMiddleware(api.SignHandler(store), jwtKey))
	http.HandleFunc("/verify", middleware.JWTMiddleware(api.VerifyHandler(store), jwtKey))

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
