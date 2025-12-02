package handler

import (
	"fmt"
	"net/http"
	// "encoding/json"
	"context"
	// "errors"
	// "math/rand"
	// "strconv"
	// "time"
	// "os"
	"log"

	// "github.com/supabase-community/supabase-go"
	// "github.com/mayo829/services/model"

	"github.com/machinebox/graphql"
)

type ResponseStruct struct {
	BlogCollection struct {
			Edges []struct {
					Node struct {
							ID          int `json:"id"`
							Name        string `json:"name"`
							Description string `json:"description"`
							CreatedAt   string `json:"createdAt"`
							UpdatedAt   string `json:"updatedAt"`
					} `json:"node"`
			} `json:"edges"`
	} `json:"blogCollection"`
}


type Blog struct{
	client *graphql.Client
}

func NewBlogHandler(client *graphql.Client) *Blog {
	return &Blog{client: client}
}

func (o *Blog) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a blog")
}

func (o *Blog) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List all blogs")
	// create a client (safe to share across requests)
	client := graphql.NewClient("https://idizvhbikkcxgfzlrmms.supabase.co/graphql/v1")

	// make a request
	req := graphql.NewRequest(`
		query{
			blogCollection(first: 1){
				edges{
					node{
						id
						name
						description
						createdAt
						updatedAt
					}
				}
			}
		}
	`)

	// set any variables
	// req.Var("key", "value")

	// set header fields
	req.Header.Set("apiKey", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImlkaXp2aGJpa2tjeGdmemxybW1zIiwicm9sZSI6ImFub24iLCJpYXQiOjE3NjQxMzQxNjQsImV4cCI6MjA3OTcxMDE2NH0.Emra9WmcIi3-jN6J0LXXUy57lXeZlYpHqp6uyjo5pb8")
	req.Header.Set("Content-Type", "application/json")

	// define a Context for the request
	ctx := context.Background()

	// run it and capture the response
	var respData ResponseStruct
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}

	// response, count, err := o.client.From("Blog").Select("*", "exact", false).ExecuteString()

	fmt.Println(respData)
}

func (o *Blog) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get a blog by ID")
}

func (o *Blog) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update a blog by ID")
}

func (o *Blog) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a blog by ID")
}