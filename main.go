package main

// GET http://localhost:8080/status

// curl -X POST http://localhost:8080/process \
//   -H "Content-Type: application/json" \
//   -d '{"data": "some important payload", "id": 12345}'
  
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// --- Structs for JSON Handling ---

// Define the expected structure of the incoming POST request body
type ProcessRequest struct {
	Data string `json:"data"`
	ID   int    `json:"id"`
}

// Define the structure of the outgoing hard-coded POST response
type ProcessResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	InputID int    `json:"received_id"`
}

func main() {
	// Heroku requires listening on the PORT environment variable.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default for local development
	}

	// Define Routes
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/process", processHandler)

	fmt.Printf("Starting server on port %s...\n", port)
	
	// Start the server
	// log.Fatal will automatically print the error and exit if the server fails to start
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// --- API Handlers ---

// GET /status: Always returns a 200 OK
func statusHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Ensure it's a GET request
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 2. Set the Content-Type header
	w.Header().Set("Content-Type", "application/json")
	
	// 3. Write a simple 200 OK response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "OK", "message": "Server is up and running"}`))
	log.Println("Served GET /status (200 OK)")
}

// POST /process: Accepts JSON, returns hard-coded JSON response
func processHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Ensure it's a POST request
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var reqBody ProcessRequest
	
	// 2. Decode the incoming JSON body
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		// Respond with a 400 Bad Request if JSON parsing fails
		http.Error(w, "Invalid JSON body: "+err.Error(), http.StatusBadRequest)
		return
	}
	
	// Log the data we received
	log.Printf("Received POST data: ID=%d, Data='%s'", reqBody.ID, reqBody.Data)

	// 3. Construct the hard-coded JSON response
	response := ProcessResponse{
		Message: "Data received successfully!",
		Status:  "processed",
		InputID: reqBody.ID, // We use the ID from the input for personalization
	}

	// 4. Set the Content-Type header and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	log.Println("Served POST /process (200 OK)")
}