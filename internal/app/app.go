package app

import (
	"log"
	"os"

	"github.com/amirtahajavadi/httpserver/internal/api"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
}

func CreateNewApp() (app *Application, err error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	workoutHandler := api.NewWorkoutHandler()
	return &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
	}, nil

}
