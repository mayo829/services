package application

import (
	"net/http"
	"log"

	"github.com/mayo829/services/handler"

	"github.com/supabase-community/supabase-go"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	client, err := NewSupabaseClient()
	if err != nil {
		log.Fatal(err)
	}

	router.Route("/blogPost", loadBlogRoutes(client))

	return router
}

func loadBlogRoutes(cli *supabase.Client, router chi.Router) {
	blogHandler := &handler.Blog{client: cli}

	router.Post("/", blogHandler.Create)
	router.Get("/", blogHandler.List)
	router.Get("/{id}", blogHandler.GetByID)
	router.Put("/{id}", blogHandler.UpdateByID)
  router.Delete("/{id}", blogHandler.DeleteByID)
}