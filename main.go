package main

import (
	"fmt"
	"context"
	"log"

	"github.com/mayo829/services/application"
	"github.com/mayo829/services/handler"
)

func main() {
	client, err := NewSupabaseClient()
	if err != nil {
		log.Fatal(err)
	}

	app := application.New(client)

  err := app.Start(context.TODO())
  if err != nil {
    fmt.Println("failed to start app:", err)
  }
}