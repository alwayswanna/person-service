package handlers

import (
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"golang.org/x/exp/slog"
	"net/http"
	"person-service/db/entity"
	"person-service/mappers"
	"person-service/model"
	"person-service/utils"
)

type PersonRepositoryImpl interface {
	SavePerson(p entity.Person) (entity.Person, error)
	DeletePerson(id uuid.UUID) (string, error)
	UpdatePerson(p entity.Person) (entity.Person, error)
	FindPerson(id *uuid.UUID) (entity.Person, error)
}

// CreatePerson godoc
// @Summary      Create new person entity
// @Description  Create new person entity
// @Tags         persons
// @Accept       json
// @Produce      json
// @Param  		 request	body    	model.PersonRequest  	true  "Model for create new person entity."
// @Success      200  		{array}   	model.PersonResponse
// @Router       /person/create [post]
func CreatePerson(logger *slog.Logger, impl PersonRepositoryImpl) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.createPerson"

		logger = logger.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req model.PersonRequest

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			logger.Error("Failed to decode message", utils.Err(err))

			render.JSON(w, r, model.Error(model.InternalError, "Error while parse request"))

			return
		}

		logger.Info("Request body decoded", slog.Any("request", req))

		entityToSave := mappers.ToPerson(req)

		savedPerson, err := impl.SavePerson(entityToSave)
		if err != nil {
			logger.Error("Error while save new person to database", utils.Err(err))

			render.JSON(w, r, model.Error(model.InternalError, "Error while save new entity"))

			return
		}

		logger.Info("Successfully save new person", slog.Any("saved_person", savedPerson))
		render.JSON(w, r, mappers.ToPersonResponse(savedPerson))
	}
}

// DeletePerson godoc
// @Summary      Delete existing persons
// @Description  Delete existing persons
// @Tags         persons
// @Accept       json
// @Produce      json
// @Param  		 id    		query    	string  					true  	"ID for remove person entity"
// @Success      200  		{array} 	model.PersonDeleteResponse
// @Router       /person/delete [delete]
func DeletePerson(logger *slog.Logger, impl PersonRepositoryImpl) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.deletePerson"

		logger = logger.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var deleteId string
		deleteId = r.URL.Query().Get("id")
		logger.Info("Request body decoded", slog.Any("entity_id", deleteId))

		id, err := impl.DeletePerson(uuid.MustParse(deleteId))
		if err != nil {
			logger.Error("Error while delete person by id", slog.String("entity_id", deleteId), utils.Err(err))

			render.JSON(w, r, model.Error(model.InternalError, fmt.Sprintf("Error while delete entity with id %s", deleteId)))

			return
		}

		logger.Info("Person with id was successfully deleted", slog.String("id", deleteId))

		render.JSON(w, r, model.CreateSuccessDeleteResponse(id))
	}
}

// UpdatePerson godoc
// @Summary      Update existing persons
// @Description  Update existing persons
// @Tags         persons
// @Accept       json
// @Produce      json
// @Param  		 request    body    	model.PersonRequest  	true  	"Model for update person entity"
// @Success      200  		{array}   	model.PersonResponse
// @Router       /person/update [put]
func UpdatePerson(logger *slog.Logger, impl PersonRepositoryImpl) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.updatePerson"

		logger = logger.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req model.PersonRequest

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			logger.Error("Failed to decode message", utils.Err(err))

			render.JSON(w, r, model.Error(model.InternalError, "Error while parse request"))

			return
		}

		logger.Info("Request body decoded", slog.Any("request", req))

		entityToSave := mappers.ToPerson(req)

		updatePerson, err := impl.UpdatePerson(entityToSave)
		if err != nil {
			logger.Error("Error while save new person to database", utils.Err(err))

			render.JSON(w, r, model.Error(model.InternalError, "Error while update entity"))

			return
		}

		logger.Info("Successfully save new person", slog.Any("updated_person", updatePerson))
		render.JSON(w, r, mappers.ToPersonResponse(updatePerson))
	}
}

// FindPerson godoc
// @Summary      Find existing persons
// @Description  Find existing persons
// @Tags         persons
// @Accept       json
// @Produce      json
// @Param		 id    query    string  				true  	"ID of person entity."
// @Success      200  {array}   model.PersonResponse
// @Router       /person/get [get]
func FindPerson(logger *slog.Logger, impl PersonRepositoryImpl) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.findPerson"

		logger = logger.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var personId string
		personId = r.URL.Query().Get("id")
		logger.Info("Request body decoded", slog.Any("entity_id", personId))

		parsedUuid := uuid.MustParse(personId)
		person, err := impl.FindPerson(&parsedUuid)
		if err != nil {
			logger.Error("Error while find person by id", slog.String("entity_id", personId), utils.Err(err))

			render.JSON(w, r, model.Error(model.InternalError, fmt.Sprintf("Error while find entity with id %s", personId)))

			return
		}

		logger.Info("Person with id was successfully found", slog.String("id", personId))

		render.JSON(w, r, mappers.ToPersonResponse(person))
	}
}
