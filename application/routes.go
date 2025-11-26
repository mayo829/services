package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/mayo829/services/handler"
)

func loadRoutes(cli *Client) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/blogPost", loadBlogRoutes(cli))

	return router
}

func loadBlogRoutes(cli *Client, router chi.Router) {
	blogHandler := &handler.Blog{client: cli}

	router.Post("/", blogHandler.Create)
	router.Get("/", blogHandler.List)
	router.Get("/{id}", blogHandler.GetByID)
	router.Put("/{id}", blogHandler.UpdateByID)
  router.Delete("/{id}", blogHandler.DeleteByID)
}