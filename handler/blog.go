package handler

import (
	"fmt"
	"net/http"
	"encoding/json"
	"errors"
	"math/rand"
	"strconv"
	"time"

	"github.com/mayo829/services/model"

	"github.com/machinebox/graphql"
	"github.com/supabase-community/supabase-go"
	"github.com/lpernett/godotenv"
)

type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create an order", r.URL.Query().Get("cursor"))
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List all orders")
}

func (o *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get an order by ID")
}

func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update an order by ID")
}

func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete an order by ID")
}