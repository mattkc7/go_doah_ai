package service

import (
	"log"
	"go_doah_ai/pkg/models" // Import the models we just defined
)

// ProcessorService represents the business logic layer for processing data.
// In a real application, this struct might hold database connections,
// configuration settings, or other dependencies.
type ProcessorService struct {
	// Example dependency:
	// db *repository.Database
}

// NewProcessorService creates and returns a new instance of the ProcessorService.
func NewProcessorService() *ProcessorService {
	return &ProcessorService{}
}

// ProcessData encapsulates the core business logic for the /process endpoint.
// It takes a request model and returns a response model. It is completely
// unaware of HTTP, making it easy to test and reuse.
func (s *ProcessorService) ProcessData(req models.ProcessRequest) models.ProcessResponse {
	// Log the incoming request to the service layer
	log.Printf("[Service Layer] Processing request for ID: %d, Data: %s", req.ID, req.Data)

	// --- Core Business Logic (currently hard-coded) ---
	// In the future, you would replace this with:
	// 1. Database lookups (calling the repository layer)
	// 2. Complex validation and calculation
	// 3. Calls to other external services

	// Hard-coded response based on the input ID
	response := models.ProcessResponse{
		Message: "Data received and processed successfully by the service layer!",
		Status:  "completed",
		InputID: req.ID,
	}

	return response
}