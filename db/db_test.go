package db

import (
	"fmt"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../.env.test")
	if err != nil {
		fmt.Println("No .env file specified")
	}
}

func TestInitDB(t *testing.T) {
	_, err := initDB()

	if err != nil {
		t.Errorf("Failed to make connection to database server")
		t.Errorf(err.Error())
	}
}

func TestGet(t *testing.T) {
	conn := Get()

	if conn != instance {
		t.Errorf("Get() does not return the same value as instance")
	}
}

func TestMigrate(t *testing.T) {
	Migrate()
}
