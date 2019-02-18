package handler

import "postastix-api/service"

var userService *service.UserService

// InitServices initalizes all services
// !!!Make sure to call this before using the handlers!!!
func InitServices() {
	userService = service.NewUserService()
}
