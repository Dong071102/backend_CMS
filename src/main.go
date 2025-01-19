package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"

	"backend_CMS_Golang/src/config"
	"backend_CMS_Golang/src/routes"
)

func main() {
	// Initialize Firebase
	config.InitializeFirebase()
	defer config.Client.Close()

	// Initialize the Router
	router := routes.SetupRoutes()

	// Set up CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"}, // Frontend URL
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	// Wrap the router with CORS middleware
	handler := c.Handler(router)

	// Start the API server
	port := ":8080"
	fmt.Printf("Server is running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
