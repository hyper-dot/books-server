package utils

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// UnmarshalJSON converts JSON byte slice to a Go struct
func ParseJSON(body io.ReadCloser, data interface{}) error {

	// Read the JSON data from the request body into a byte slice
	err := json.NewDecoder(body).Decode(data)
	if err != nil {
		return err
	}
	return nil
}

// RESPOND WITH JSON
func RespondWithJSON(payload interface{}, w http.ResponseWriter, status int) {
	result, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Faile to Marshall payload: %s", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(result)
}

// RESPOND WITH ERR RESPONSE
func RespondWithError(message string, w http.ResponseWriter, status int, ctx context.Context) {
	RespondWithJSON(ErrorResponse{Error: message}, w, status)
	ctx.Done()
}
