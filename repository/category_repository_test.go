package repository

import (
	"fmt"
	"postastix-api/model"
	"testing"

	"github.com/joho/godotenv"
)

var _mockCategoryID uint
var _mockCategoryName string

func init() {
	err := godotenv.Load("../.env.test")

	if err != nil {
		fmt.Println("No .env file was specified")
	}
}

func TestPush(t *testing.T) {
	_mockCategoryName := _testFaker.Person().Name()

	newCategory := model.Category{
		Name: _mockCategoryName,
	}

	_testCategoryRepo.Push(&newCategory)
	_mockUserID = newCategory.ID
}

func TestCategoryGet(t *testing.T) {
	categories := _testCategoryRepo.Get()

	if len(categories) == 0 {
		t.Errorf("Get() doesn't return any data")
	}
}

func TestCategoryFind(t *testing.T) {
	category := _testCategoryRepo.Find(_mockCategoryID)

	if category.Name != _mockCategoryName {
		t.Errorf("Find() doesn't return any data")
	}
}

func TestCategoryDelete(t *testing.T) {
	_testCategoryRepo.Delete(_mockCategoryID)
	category := _testCategoryRepo.Find(_mockCategoryID)

	if category.ID != 0 {
		t.Errorf("Delete() failed to delete data")
	}
}
