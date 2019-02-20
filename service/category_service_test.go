package service

import "testing"

var _mockCategoryID uint
var _mockName string

func TestCategoryCreate(t *testing.T) {
	_mockName = _testFaker.Company().JobTitle()
	category, err := _testCategoryService.Create(_mockName)

	if err != nil {
		t.Errorf(err.Error())
	}

	if category.Name != _mockName {
		t.Errorf("Something error with Create()")
	}

	_mockCategoryID = category.ID
}

func TestCategoryGet(t *testing.T) {
	category := _testCategoryService.Get()

	if category == nil {
		t.Errorf("Something error with Get()")
	}
}

func TestCategoryFind(t *testing.T) {
	category, err := _testCategoryService.Find(_mockCategoryID)

	if err != nil {
		t.Errorf(err.Error())
	}

	if category.ID != _mockCategoryID {
		t.Errorf("Something wrong with Find()")
	}
}

func TestCategoryUpdate(t *testing.T) {
	_mockNewName := _testFaker.Person().Name()
	updatedUser, err := _testCategoryService.Update(_mockCategoryID, _mockNewName)

	if err != nil {
		t.Errorf(err.Error())
	}

	if updatedUser.Name != _mockNewName {
		t.Errorf("Something wrong with Update()")
	}
}

func TestCategoryDelete(t *testing.T) {
	err := _testCategoryService.Delete(_mockCategoryID)

	if err != nil {
		t.Errorf(err.Error())
	}
}
