package handlers

import (
	"encoding/json"
	"net/http"
)

func HandleGet(w http.ResponseWriter, r *http.Request) error {
	response := map[string]string{
		"message": "Hello, World!",
	}
	return json.NewEncoder(w).Encode(response)
}
