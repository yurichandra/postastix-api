package handler

import "github.com/dewadg/postastix-api/service"

var userService *service.UserService

// Init sets all the required services
func Init() {
	userService = service.GetUserService()
}
