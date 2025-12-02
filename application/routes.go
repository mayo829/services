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

	client, err := handler.NewSupabaseClient()
	if err != nil {
		log.Fatal(err)
	}

	router.Route("/blogpost", func(r chi.Router) {
		loadBlogRoutes(client, r)
	})

	return router
}

func loadBlogRoutes(cli *graphql.Client, router chi.Router) {
	blogHandler := handler.NewBlogHandler(cli)

	router.Post("/", blogHandler.Create)
	router.Get("/", blogHandler.List)
	router.Get("/{id}", blogHandler.GetByID)
	router.Put("/{id}", blogHandler.UpdateByID)
  router.Delete("/{id}", blogHandler.DeleteByID)
}