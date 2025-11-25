package api

import (
	"encoding/json"
	"log"
	"net/http"

	"go_doah_ai/internal/service" // Import the business logic layer
	"go_doah_ai/pkg/models"       // Import the data structures
)

// Handler holds the service layer dependencies required by the handlers.
// This is how you "inject" the service layer into the API layer.
type Handler struct {
	Processor *service.ProcessorService
}

// NewHandler initializes a new Handler with the given service.
func NewHandler(processorService *service.ProcessorService) *Handler {
	return &Handler{
		Processor: processorService,
	}
}

// StatusHandler (GET /status) always returns a 200 OK.
func (h *Handler) StatusHandler(w http.ResponseWriter, r *http.Request) {
	// 1. HTTP/Web check: Ensure correct method
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 2. HTTP/Web response formatting
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "OK", "message": "API Layer is healthy and connected"}`))
	log.Println("Served GET /status (200 OK)")
}

// ProcessHandler (POST /process) handles decoding the request, calling the service,
// and encoding the response.
func (h *Handler) ProcessHandler(w http.ResponseWriter, r *http.Request) {
	// 1. HTTP/Web check: Ensure correct method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 2. HTTP/Web input decoding
	var reqBody models.ProcessRequest
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid JSON body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// 3. Delegation to Service Layer (Core Principle: API layer doesn't do logic)
	// We pass the clean request model to the business layer.
	response := h.Processor.ProcessData(reqBody)

	// 4. HTTP/Web output encoding
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

	log.Printf("Served POST /process (200 OK) for ID: %d", reqBody.ID)
}