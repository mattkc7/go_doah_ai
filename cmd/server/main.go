package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go_doah_ai/internal/api"
	"go_doah_ai/internal/service"
)

// go run ./cmd/server/main.go

func main() {
	// 1. Initialize Service Layer (Business Logic)
	// This is the cleanest layer; it has no knowledge of HTTP.
	processorService := service.NewProcessorService()

	// 2. Initialize API Layer (HTTP Controllers)
	// This layer requires the Service Layer to do its work.
	handler := api.NewHandler(processorService)

	// 3. Configure HTTP Server and Routing
	http.HandleFunc("/status", handler.StatusHandler)
	http.HandleFunc("/process", handler.ProcessHandler)

	// 4. Start the Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Starting refactored server on port %s...\n", port)
	
	// log.Fatal will automatically print the error and exit if the server fails to start
	log.Fatal(http.ListenAndServe(":"+port, nil))
}