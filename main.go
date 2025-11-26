package main

import (
	"fmt"
	// "context"
	"log"

	"github.com/mayo829/services/application"
)

func main() {
	app := application.New()

  err := app.Start()
  if err != nil {
    fmt.Println("failed to start app:", err)
  }
}