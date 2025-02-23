package di

import (
	"log"
	"net/http"
	"tiny_files/config"
	"tiny_files/routes"
)

func Run() {
	config.LoadConfig()
	// Initialize MinIO client
	minioClient := InitializeMinioClient()
	// Set up routes
	routes.SetupRoutes(minioClient)
	log.Printf("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
