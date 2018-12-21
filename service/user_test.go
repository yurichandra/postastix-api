package service

import (
	"fmt"
	"testing"

	"github.com/joho/godotenv"
)

var userSrv *UserService

func init() {
	err := godotenv.Load("../.env.test")
	if err != nil {
		fmt.Println("No .env file specified")
	}

	userSrv = GetUserService()
}

func TestUserCreate(t *testing.T) {
	newUser, err := userSrv.Create("johndoe", "John Doe", "doejohn")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if newUser.Name != "johndoe" || newUser.FullName != "John Doe" {
		t.Errorf("Create() does not store user well")
	}
}

func TestUserUpdate(t *testing.T) {
	updatedUser, err := userSrv.Update(4, "johnbeep", "John Beep")
	if err != nil {
		t.Errorf(err.Error())
	}

	if updatedUser.Name != "johnbeep" || updatedUser.FullName != "John Beep" {
		t.Errorf("Update() does not update user well")
	}
}

func TestUserDelete(t *testing.T) {
	err := userSrv.Delete(5)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestUserChangePassword(t *testing.T) {
	err := userSrv.ChangePassword(7, "dudu", "dada", "dada")
	if err != nil {
		t.Errorf(err.Error())
	}
}
