package object

import (
	"errors"
	"net/http"
)

// StoreUserRequest represents request object for
// creating a user.
type StoreUserRequest struct {
	Name            string `json:"name"`
	FullName        string `json:"fullName"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

// Bind checks the request.
func (req *StoreUserRequest) Bind(r *http.Request) error {
	if req.Name == "" {
		return errors.New("Field `name` cannot be empty")
	}
	if req.FullName == "" {
		return errors.New("Field `fullName` cannot be empty")
	}
	if req.Password == "" {
		return errors.New("Field `password` cannot be empty")
	}
	if req.ConfirmPassword == "" {
		return errors.New("Field `confirmPassword` cannot be empty")
	}
	if req.Password != req.ConfirmPassword {
		return errors.New("Field `password` and `confirmPassword` not match")
	}
	return nil
}
