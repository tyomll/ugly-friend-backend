package utils

import (
	"encoding/json"
	"net/http"
)

func HandleRouterBodyRequest(w http.ResponseWriter, r *http.Request, requestBody interface{}, updateFunc func(interface{}) (interface{}, error)) {
	err := json.NewDecoder(r.Body).Decode(requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	res, err := updateFunc(requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
