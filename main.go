package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file specified")
	}

	switch command() {
	case "serve":
		router := createRouter()

		fmt.Printf("App running on port %s\n", os.Getenv("APP_PORT"))
		http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), router)
		break
	default:
		fmt.Println("Invalid command")
	}
}

func command() string {
	args := os.Args[1:]

	if len(args) > 0 {
		return args[0]
	}
	return ""
}
