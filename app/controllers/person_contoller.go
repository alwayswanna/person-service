package controllers

import (
	"github.com/go-chi/chi/v5"
	"golang.org/x/exp/slog"
	"person-service/db/repository"
	"person-service/handlers"
)

func RegisterPersonHandlers(logger *slog.Logger, router *chi.Mux, storage *repository.PersonRepositoryImpl) {
	router.Post("/api/v1/person/create", handlers.CreatePerson(logger, storage))
	router.Delete("/api/v1/person/delete", handlers.DeletePerson(logger, storage))
	router.Put("/api/v1/person/update", handlers.UpdatePerson(logger, storage))
	router.Get("/api/v1/person/get/id", handlers.FindPersonById(logger, storage))
	router.Get("/api/v1/persons", handlers.LoadPersons(logger, storage))
	router.Get("/api/v1/person/get/login", handlers.FindPersonByLogin(logger, storage))
}
