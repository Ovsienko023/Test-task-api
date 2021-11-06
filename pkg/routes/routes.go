package routes

import (
	"user_api/pkg/handler"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux) {
	r.Mount("/api/v1", r)

	r.Get("/", handler.EchoData)

	r.Get("/users", handler.SearchUsers)
	r.Post("/users", handler.CreateUser)
	r.Get("/users/{user_id}", handler.GetUser)
	r.Patch("/users/{id}", handler.UpdateUser)
	r.Delete("/users/{user_id}", handler.DeleteUser)

}
