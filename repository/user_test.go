package repository

import (
	"fmt"
	"postastix-api/model"
	"testing"

	"github.com/joho/godotenv"
)

var _testUserRepo *UserRepository
var _testUserID uint

func init() {
	err := godotenv.Load("../.env.test")
	if err != nil {
		fmt.Println("No .env file specified")
	}

	_testUserRepo = NewUserRepository()
}

func TestUserGet(t *testing.T) {
	_testUserRepo.Get()
}

func TestUserPush(t *testing.T) {
	newUser := model.User{
		Name:     "johndoe",
		FullName: "John Doe",
		Password: "doejohn",
	}

	_testUserRepo.Push(&newUser)
	_testUserID = newUser.ID
}

func TestUserFind(t *testing.T) {
	user := _testUserRepo.Find(_testUserID)

	if user.Name != "johndoe" {
		t.Errorf("Find() returns wrong user.")
	}
}

func TestUserDelete(t *testing.T) {
	_testUserRepo.Delete(_testUserID)
	user := _testUserRepo.Find(_testUserID)

	if user.ID != 0 {
		t.Errorf("Delete() failed to remove user")
	}
}
