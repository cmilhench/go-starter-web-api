package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Home provides a simple home endpoint
func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "Welcome to the Go Web Server!")
}

// Health provides a simple health endpoint
func Health(w http.ResponseWriter, r *http.Request) {
	status := map[string]string{
		"status":  "healthy",
		"version": "1.0.0",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}
