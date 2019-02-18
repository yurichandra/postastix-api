package main

import (
	"fmt"
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
		serveHTTP()
		break
	case "migrate":
		runMigration()
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
