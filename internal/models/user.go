package models

import (
	"fmt"
	"time"
)

// User represents a simple user model
type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// Validate performs basic validation on the user
func (u *User) Validate() error {
	if u.Username == "" {
		return fmt.Errorf("username is required")
	}
	if u.Email == "" {
		return fmt.Errorf("email is required")
	}
	return nil
}
