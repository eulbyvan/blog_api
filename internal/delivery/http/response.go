package http

import (
	"encoding/json"
	"net/http"
)

// custom response
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// json response sender
func JSONResponse(w http.ResponseWriter, statusCode int, status string, message string, data interface{}) {
	response := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
