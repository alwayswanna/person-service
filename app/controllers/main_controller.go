package controllers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"
	"person-service/utils"
)

func RegisterMiddlewareHandlers(logger *slog.Logger, router *chi.Mux) {
	/* register middleware filters */
	router.Use(middleware.RequestID)
	router.Use(utils.New(logger))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
}
