package conf

import (
	"github.com/Mario-Benedict/note-api/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes(router *chi.Mux) {
	router.Use(middleware.Logger)
	router.Use(middleware.AllowContentType("application/json"))

	healthController := controllers.HealthController{}
	noteController := controllers.NoteController{}

	router.Get("/", healthController.Check)
	router.Get("/notes", noteController.GetAll)
	router.Get("/notes/{id}", noteController.GetByID)
	router.Post("/notes", noteController.Create)
	router.Put("/notes", noteController.Update)
	router.Delete("/notes", noteController.Delete)
}
