package main

import (
	"flag"
	"fmt"
	"go-backend-server-template/internal/app"
	"go-backend-server-template/internal/routes"
	"net/http"
	"time"
)

func main() {
	var port int

	// take input from command line if not present fall back to 8080
	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse()

	// Initialize new application
	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	// defer the closing of database connection
	defer app.DB.Close()

	// set up routes
	r := routes.SetUpRoutes(*app)

	// server properties
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler:      r,
	}

	app.Logger.Printf("Started go server at %d", port)

	// Start the server
	err = server.ListenAndServe()

	if err != nil {
		app.Logger.Fatal(err)
	}
}
