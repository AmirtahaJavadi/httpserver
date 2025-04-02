package main

import (
	"fmt"

	"github.com/amirtahajavadi/httpserver/internal/app"
)

func main() {
	app, err := app.CreateNewApp()
	if err != nil {
		panic(err)
	}
	fmt.Println("hello world", app.Logger)

}
