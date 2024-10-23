package controllers

import (
	"crypto/rsa"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"golang.org/x/exp/slog"
	_ "person-service/docs"
	"person-service/handlers"
	"person-service/utils"
)

func RegisterMiddlewareHandlers(logger *slog.Logger, router *chi.Mux, rsaPubKey *rsa.PublicKey) {
	/* register middleware filters */
	handlers.Init(rsaPubKey)
	router.Use(middleware.RequestID)
	if rsaPubKey != nil {
		router.Use(handlers.JwtBearerValidation)
	}
	router.Use(utils.New(logger))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:9902/swagger/doc.json"), //The url pointing to API definition
	))
}
