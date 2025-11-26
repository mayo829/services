package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/mayo829/services/application"
)

func main() {
	app := application.New()

  err := app.Start(context.TODO())
  if err != nil {
    fmt.Println("failed to start app:", err)
  }
}