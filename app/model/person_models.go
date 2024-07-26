package model

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

// PersonRequest model info
// @Description Model for create or update person entity.
type PersonRequest struct {
	Id        uuid.UUID `json:"id,omitempty"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Age       int       `json:"age"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

// PersonResponse model info
// @Description Model for response on API operations.
type PersonResponse struct {
	Id        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Age       int       `json:"age"`
	Timestamp time.Time `json:"timestamp"`
}

// PersonDeleteResponse model info
// @Description Model for response on delete operation.
type PersonDeleteResponse struct {
	Message string `json:"message"`
}

func CreateSuccessDeleteResponse(id string) PersonDeleteResponse {
	return PersonDeleteResponse{
		Message: fmt.Sprintf("Person with id: %s was successfully deleted", id),
	}
}
