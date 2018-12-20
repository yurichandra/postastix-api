package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // MySQL driver
)

var instance *gorm.DB

func initDB() (*gorm.DB, error) {
	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	return gorm.Open("mysql", connString)
}

// Get returns DB connection singleton.
func Get() *gorm.DB {
	if instance == nil {
		conn, err := initDB()
		if err != nil {
			panic(err.Error())
		}

		instance = conn
	}
	return instance
}

// Migrate creates tables for available models.
func Migrate() {
	Get().AutoMigrate(
		&User{},
		&Category{},
		&Post{},
	)
}
