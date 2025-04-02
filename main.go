package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/amirtahajavadi/httpserver/internal/app"
)

func main() {
	port := flag.Int("port", 5000, "Server port")
	flag.Parse()
	app, err := app.CreateNewApp()
	if err != nil {
		panic(err)
	}
	app.Logger.Println("app is running")

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", *port),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	http.HandleFunc("/health", HealthChecker)

	err = server.ListenAndServe()
	if err != nil {
		return
	}

}

func HealthChecker(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "status is available\n")
}
