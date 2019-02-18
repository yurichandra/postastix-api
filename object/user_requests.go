package object

import "errors"

// StoreUserRequest represents request object for
// creating a user.
type StoreUserRequest struct {
	Name            string `json:"name"`
	FullName        string `json:"fullName"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

// Validate request.
func (r *StoreUserRequest) Validate() error {
	if r.Name == "" {
		return errors.New("Property `name` cannot be empty")
	}
	if r.FullName == "" {
		return errors.New("Property `fullName` cannot be empty")
	}
	if r.Password == "" {
		return errors.New("Property `password` cannot be empty")
	}
	if r.ConfirmPassword == "" {
		return errors.New("Property `confirmPassword` cannot be empty")
	}
	if r.Password != r.ConfirmPassword {
		return errors.New("Property `password` and `confirmPassword` not match")
	}
	return nil
}
