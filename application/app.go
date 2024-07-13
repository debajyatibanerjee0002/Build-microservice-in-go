package application

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	rdb    *redis.Client
}

func New() *App {
	// Initialize the Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Adjust the address as needed
		// Add any other necessary options, such as Password and DB
	})

	app := &App{
		router: loadRoutes(),
		rdb:    rdb,
	}

	return app
}

// receiver function
func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	err := a.rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to connect to redis: %w ", err)
	}

	err = server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to listen to server: %w ", err)
	}
	return nil
}
