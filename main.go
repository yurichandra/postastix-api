package main

import (
	"fmt"
	"os"

	"github.com/dewadg/postastix-api/db"
	"github.com/dewadg/postastix-api/handler"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file specified")
	}

	app := createApp(db.Get(), createRouter())

	switch command() {
	case "serve":
		handler.Init()
		app.Run()
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
