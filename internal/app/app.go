package app

import (
	"log"
	"os"
)

type Application struct {
	Logger *log.Logger
}

func CreateNewApp() (app *Application, err error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	return &Application{
		Logger: logger,
	}, nil
}
