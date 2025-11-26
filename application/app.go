package application

import (
  "context"
  "fmt"
  "net/http"
)

type App struct {
  router http.Handler
	client *Client
}

func New(client *Client) *App {
  app := &App{
    router: loadRoutes(client),
		client: client,
  }
  
  return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	err := server.ListenAndServe()
	if err != nil {
    return fmt.Errorf("failed to start server: %w", err)
	}
  
  return nil
}