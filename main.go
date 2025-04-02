package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/amirtahajavadi/httpserver/internal/app"
	"github.com/amirtahajavadi/httpserver/internal/routes"
)

func main() {
	port := flag.Int("port", 5000, "Server port")
	flag.Parse()
	app, err := app.CreateNewApp()
	if err != nil {
		panic(err)
	}
	app.Logger.Println("app is running")
	r := routes.SetupRoutes(app)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", *port),
		IdleTimeout:  time.Minute,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		return
	}

}
