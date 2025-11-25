package models

// --- Data Structures for Request/Response ---

// ProcessRequest defines the expected structure of the incoming JSON body
// for the POST /process endpoint.
type ProcessRequest struct {
	Data string `json:"data"`
	ID   int    `json:"id"`
}

// ProcessResponse defines the structure of the JSON response returned
// by the POST /process endpoint.
type ProcessResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	InputID int    `json:"received_id"`
}

// --- Data Structures for Internal Use (Example) ---

// InternalUser is an example of a type that might be used internally
// by the service or repository layer, perhaps mirroring a database record.
type InternalUser struct {
	UserID    int
	InputData string
	Timestamp int64
}