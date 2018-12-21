package service

import (
	"github.com/dewadg/postastix-api/repository"
)

var user *UserService

// GetUserService returns user service singleton.
func GetUserService() *UserService {
	if user == nil {
		user = &UserService{
			repo: repository.GetUserRepository(),
		}
	}
	return user
}
