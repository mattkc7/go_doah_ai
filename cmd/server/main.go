package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"go_doah_ai/internal/api"
	"go_doah_ai/internal/service"
)

// go run ./cmd/server/main.go

const staticDir = "frontend"
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
	http.HandleFunc("/", spaHandler)
	
	// 4. Start the Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Starting refactored server on port %s...\n", port)
	
	// log.Fatal will automatically print the error and exit if the server fails to start
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// spaHandler (Single Page Application Handler)
// This function is the core of embedding the React frontend into the Go server.
// It serves static files normally but serves index.html for unknown paths,
// allowing the React router to handle client-side URLs (like /dashboard).
func spaHandler(w http.ResponseWriter, r *http.Request) {
	// Use Go's built-in file server to serve content from the 'frontend' directory
	// This is a powerful part of Go's standard library.
	fs := http.FileServer(http.Dir(staticDir))

	// Check if the requested path exists as a physical file on disk (e.g., /bundle.js).
	if _, err := os.Stat(filepath.Join(staticDir, r.URL.Path)); err == nil {
		// If the file exists, serve it directly.
		fs.ServeHTTP(w, r)
		return
	}

	// If the file does NOT exist (e.g., the user navigated to /dashboard,
	// which is a client-side route), we MUST serve the root index.html.
	// The React application inside index.html then takes over and renders the correct view.
	log.Printf("Path not found: %s. Serving index.html for SPA routing.", r.URL.Path)
	http.ServeFile(w, r, filepath.Join(staticDir, "index.html"))
}
