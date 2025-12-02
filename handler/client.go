package handler

import (
	"fmt"
	"os"

	"github.com/supabase-community/supabase-go"
	"github.com/joho/godotenv"
)

func NewSupabaseClient() (*graphql.Client, error) {
	_ = godotenv.Load() // don't fatal—allow environment-only deployments

	// url := os.Getenv("SUPABASE_PROJECT_URL")
	project_id := os.Getenv("SUPABASE_PROJECT_ID")
	// key := os.Getenv("SUPABASE_API_KEY")

	if project_id == "" {
		return nil, fmt.Errorf("missing Supabase credentials")
	}

	return graphql.NewClient("https://{project_id}.supabase.co/graphql/v1")
}
