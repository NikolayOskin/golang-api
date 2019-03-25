package handler

import (
	"encoding/json"
	"net/http"
)

// ValidationError is a validator error struct
type ValidationError struct {
	Error map[string]interface{} `json:"error"`
}

// RespondJSON - respond json with status
func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// RespondError - respond error json with error
func RespondError(w http.ResponseWriter, code int, message string) {
	RespondJSON(w, code, map[string]string{"error": message})
}

// RespondValidationError - respond validation errors json
func RespondValidationError(w http.ResponseWriter, code int, message map[string]interface{}) {
	result := &ValidationError{message}
	response, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-type", "applciation/json")
	w.Write([]byte(response))
}
