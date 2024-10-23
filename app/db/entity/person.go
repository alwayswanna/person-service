package entity

import (
	"github.com/google/uuid"
	"time"
)

type Person struct {
	Id        *uuid.UUID
	Login     string
	FirstName string
	LastName  string
	Age       int
	Timestamp *time.Time
}
