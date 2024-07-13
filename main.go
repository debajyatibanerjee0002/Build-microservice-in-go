package main

import (
	"context"
	"fmt"

	"github.com/debajyatibanerjee0002/Build-microservice-in-go.git/application"
)

// func main() {
// 	// initialize the router
// 	router := chi.NewRouter()
// 	// for getting log values
// 	router.Use(middleware.Logger)
// 	// for handling HTTP requests
// 	router.Get("/hello", basicHandler)

// 	server := &http.Server{
// 		Addr:    ":3000",
// 		Handler: router,
// 	}
// 	// to listen the server
// 	err := server.ListenAndServe()
// 	if err != nil {
// 		fmt.Println("failed to listen to server ", err)
// 	}
// }

// func basicHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("hello world"))
// }

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start application: ", err)
	}
}
