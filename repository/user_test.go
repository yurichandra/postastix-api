package repository

import (
	"fmt"
	"postastix-api/model"
	"testing"

	"github.com/joho/godotenv"
)

var _mockUserID uint
var _mockUserName string
var _mockUserFullName string
var _mockUserPassword string

func init() {
	err := godotenv.Load("../.env.test")
	if err != nil {
		fmt.Println("No .env file specified")
	}

	_testUserRepo = NewUserRepository()
}

func TestUserPush(t *testing.T) {
	_mockUserName = _testFaker.Person().FirstName()
	_mockUserFullName = _testFaker.Person().FirstNameMale()
	_mockUserPassword = _testFaker.Lorem().Word()

	newUser := model.User{
		Name:     _mockUserName,
		FullName: _mockUserFullName,
		Password: _mockUserPassword,
	}

	_testUserRepo.Push(&newUser)
	_mockUserID = newUser.ID
}

func TestUserGet(t *testing.T) {
	userList := _testUserRepo.Get()

	if len(userList) == 0 {
		t.Errorf("Get() does not return any user")
	}
}

func TestUserFind(t *testing.T) {
	user := _testUserRepo.Find(_mockUserID)

	if user.Name != _mockUserName {
		t.Errorf("Find() returns wrong user.")
	}
}

func TestUserDelete(t *testing.T) {
	_testUserRepo.Delete(_mockUserID)
	user := _testUserRepo.Find(_mockUserID)

	if user.ID != 0 {
		t.Errorf("Delete() failed to remove user")
	}
}
