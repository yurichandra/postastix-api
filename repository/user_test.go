package repository

import (
	"fmt"
	"testing"

	"github.com/dewadg/postastix-api/db"
	"github.com/joho/godotenv"
)

var userRepo *UserRepository

func init() {
	err := godotenv.Load("../.env.test")
	if err != nil {
		fmt.Println("No .env file specified")
	}

	userRepo = new(UserRepository)
}

func TestUserGet(t *testing.T) {
	userRepo.Get()
}

func TestUserPush(t *testing.T) {
	newUser := &db.User{
		Name:     "johndoe",
		FullName: "John Doe",
		Password: "doejohn",
	}

	userRepo.Push(newUser)
}

func TestUserFind(t *testing.T) {
	user := userRepo.Find(1)

	if user.Name != "dewadg" {
		t.Errorf("Find() returns wrong user.")
	}
}

func TestUserDelete(t *testing.T) {
	userRepo.Delete(2)
	user := userRepo.Find(2)

	if user.ID != 0 {
		t.Errorf("Delete() failed to remove user")
	}
}
