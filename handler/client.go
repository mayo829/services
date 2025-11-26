package handler

import (
	"fmt"
	"os"

	"github.com/supabase-community/supabase-go"
	"github.com/joho/godotenv"
)

func NewSupabaseClient() (*supabase.Client, error) {
	_ = godotenv.Load() // don't fatal—allow environment-only deployments

	url := os.Getenv("SUPABASE_PROJECT_URL")
	key := os.Getenv("SUPABASE_API_KEY")

	if url == "" || key == "" {
		return nil, fmt.Errorf("missing Supabase credentials")
	}

	return supabase.NewClient(url, key, nil)
}
