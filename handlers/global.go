package handlers

import (
	"encoding/json"
	"net/http"
)

func HandleGet(w http.ResponseWriter, _ *http.Request) error {
	response := map[string]string{
		"message": "Hello, World!",
	}
	return json.NewEncoder(w).Encode(response)
}

func HandlePost(w http.ResponseWriter, r *http.Request) error {
	var body map[string]any
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(body)
}
