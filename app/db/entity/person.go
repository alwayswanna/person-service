package entity

import (
	"github.com/google/uuid"
	"time"
)

type Person struct {
	Id        *uuid.UUID
	FirstName string
	LastName  string
	Age       int
	Timestamp *time.Time
}
