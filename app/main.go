package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"golang.org/x/exp/slog"
	"net/http"
	"os"
	"person-service/config"
	"person-service/controllers"
	"person-service/db/repository"
	"person-service/utils"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

var logger *slog.Logger
var configuration *config.Config
var storage *repository.PersonRepositoryImpl
var router *chi.Mux

func init() {
	/* init configuration */
	configuration = config.LoadConfiguration()
	/* init logger */
	logger = setupLogger(configuration.Env)
	/* init database */
	db, err := repository.New(configuration.Datasource)
	if err != nil {
		logger.Error("Failed while init database connection", utils.Err(err))
		os.Exit(1)
	}
	storage = db

	/* init router */
	router = chi.NewRouter()
	controllers.RegisterCorsMiddlewareHandlers(router)

	/* init security | mock security for integration testing */
	if configuration.Security.Module != "" {
		rsaPubKey, rsaErr := utils.ConstructRsaPublicKey(configuration.Security.Module, configuration.Security.Exponent)
		if rsaErr != nil {
			logger.Error("Failed to create rsa.PublicKey", utils.Err(rsaErr))
			os.Exit(1)
		}
		controllers.RegisterMiddlewareHandlers(logger, router, rsaPubKey)
	} else {
		controllers.RegisterMiddlewareHandlers(logger, router, nil)
	}

	/* register api handlers */
	controllers.RegisterPersonHandlers(logger, router, storage)
}

// @title           person-service API
// @version         1.0
// @description     This is a sample server on go-lang.
// @contact.name    API Support
// @contact.email   support@swagger.io
// @host      		localhost:9902
// @BasePath  		/api/v1
// @externalDocs.description  API for create/update/delete/edit persons.
func main() {
	logger.Info("Starting person-service ... ", slog.String("env", configuration.Env))

	logger.Info("Starting http-s: ", slog.Int("port", configuration.Server.Port))

	server := setupHttpServer(configuration, router)

	if err := server.ListenAndServe(); err != nil {
		logger.Error("Http-s start failed, ", utils.Err(err))
	}

	logger.Error("Http-s stopped.")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}

func setupHttpServer(configuration *config.Config, router *chi.Mux) *http.Server {
	return &http.Server{
		Addr:         fmt.Sprintf("localhost:%d", configuration.Server.Port),
		Handler:      router,
		IdleTimeout:  configuration.Server.IdleTimeout,
		ReadTimeout:  configuration.Server.Timeout,
		WriteTimeout: configuration.Server.Timeout,
	}
}
