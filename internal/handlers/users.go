package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/cmilhench/go-starter-web-api/internal/models"
)

// UserStore simulates an in-memory database
type UserStore struct {
	mu    sync.RWMutex
	users map[string]models.User
}

// NewUserStore creates a new UserStore
func NewUserStore() *UserStore {
	return &UserStore{
		users: make(map[string]models.User),
	}
}

// Global user store
var store = NewUserStore()

// Helper function to send JSON response
func sendJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// GetUsers handles GET requests to retrieve users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	store.mu.RLock()
	defer store.mu.RUnlock()

	users := make([]models.User, 0, len(store.users))
	for _, user := range store.users {
		users = append(users, user)
	}

	sendJSONResponse(w, http.StatusOK, users)
}

// GetUser handles GET requests for a specific user
func GetUser(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from URL
	parts := strings.Split(r.URL.Path, "/")
	userID := parts[len(parts)-1]

	store.mu.RLock()
	defer store.mu.RUnlock()

	user, exists := store.users[userID]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	sendJSONResponse(w, http.StatusOK, user)
}

// CreateUser handles POST requests to create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Read request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Unmarshal JSON
	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate user
	if err := user.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate unique ID
	user.ID = fmt.Sprintf("%d", time.Now().UnixNano())
	user.CreatedAt = time.Now()

	// Store user
	store.mu.Lock()
	store.users[user.ID] = user
	store.mu.Unlock()

	// Respond with created user
	sendJSONResponse(w, http.StatusCreated, user)
}

// UpdateUser handles PUT requests to update a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from URL
	parts := strings.Split(r.URL.Path, "/")
	userID := parts[len(parts)-1]

	// Read request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Unmarshal JSON
	var updatedUser models.User
	if err := json.Unmarshal(body, &updatedUser); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	store.mu.Lock()
	defer store.mu.Unlock()

	// Check if user exists
	existingUser, exists := store.users[userID]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Update user fields
	if updatedUser.Username != "" {
		existingUser.Username = updatedUser.Username
	}
	if updatedUser.Email != "" {
		existingUser.Email = updatedUser.Email
	}

	// Store updated user
	store.users[userID] = existingUser

	sendJSONResponse(w, http.StatusOK, existingUser)
}

// DeleteUser handles DELETE requests to remove a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from URL
	parts := strings.Split(r.URL.Path, "/")
	userID := parts[len(parts)-1]

	store.mu.Lock()
	defer store.mu.Unlock()

	// Check if user exists
	_, exists := store.users[userID]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Delete user
	delete(store.users, userID)

	w.WriteHeader(http.StatusNoContent)
}
