package model

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type PersonRequest struct {
	Id        uuid.UUID `json:"id,omitempty"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Age       int       `json:"age"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

type PersonResponse struct {
	Id        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Age       int       `json:"age"`
	Timestamp time.Time `json:"timestamp"`
}

type PersonDeleteResponse struct {
	Message string `json:"message"`
}

func CreateSuccessDeleteResponse(id string) PersonDeleteResponse {
	return PersonDeleteResponse{
		Message: fmt.Sprintf("Person with id: %s was successfully deleted", id),
	}
}
