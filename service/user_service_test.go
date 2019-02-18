package service

import (
	"testing"
)

var _mockUserID uint
var _mockUserName string
var _mockUserFullName string
var _mockUserPassword string

func TestUserCreate(t *testing.T) {
	_mockUserName = _testFaker.Person().Name()
	_mockUserFullName = _testFaker.Person().FirstNameMale()
	_mockUserPassword = _testFaker.Lorem().Word()

	newUser, err := _testUserService.Create(_mockUserName, _mockUserFullName, _mockUserPassword)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if newUser.Name != _mockUserName || newUser.FullName != _mockUserFullName {
		t.Errorf("Create() does not store user well")
	}

	_mockUserID = newUser.ID
}

func TestUserChangePassword(t *testing.T) {
	_mockOldPassword := _mockUserPassword
	_mockUserPassword = _testFaker.Lorem().Word()

	err := _testUserService.ChangePassword(_mockUserID, _mockOldPassword, _mockUserPassword, _mockUserPassword)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestUserFind(t *testing.T) {
	user, err := _testUserService.Find(_mockUserID)
	if err != nil {
		t.Errorf(err.Error())
	}

	if user.ID != _mockUserID {
		t.Errorf("Find() does not return the specified user")
	}
}

func TestUserUpdate(t *testing.T) {
	_mockUserName = _testFaker.Person().Name()
	_mockUserFullName = _testFaker.Person().FirstNameMale()

	updatedUser, err := _testUserService.Update(_mockUserID, _mockUserName, _mockUserFullName)
	if err != nil {
		t.Errorf(err.Error())
	}

	if updatedUser.Name != _mockUserName || updatedUser.FullName != _mockUserFullName {
		t.Errorf("Update() does not update user well")
	}
}

func TestUserDelete(t *testing.T) {
	err := _testUserService.Delete(_mockUserID)
	if err != nil {
		t.Errorf(err.Error())
	}
}
