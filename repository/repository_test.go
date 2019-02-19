package repository

import (
	"fmt"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/joho/godotenv"
)

var _testFaker faker.Faker
var _testUserRepo *UserRepository
var _testCategoryRepo *CategoryRepository

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env.test")
	if err != nil {
		fmt.Println("No .env file specified")
	}

	_testFaker = faker.New()
	_testUserRepo = NewUserRepository()
	_testCategoryRepo = NewCategoryRepository()

	m.Run()
}
