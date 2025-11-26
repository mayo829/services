package handler

import (
	"fmt"
	"net/http"
	"encoding/json"
	"errors"
	"math/rand"
	"strconv"
	"time"
	"os"
	"log"

	"github.com/mayo829/services/model"

	"github.com/machinebox/graphql"
	"github.com/supabase-community/supabase-go"
)

type Blog struct{
	client *Client
}

func (o *Blog) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a blog")
}

func (o *Blog) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List all blogs")
	data, count, err := o.client.From("countries").Select("*", "exact", false).Execute()
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