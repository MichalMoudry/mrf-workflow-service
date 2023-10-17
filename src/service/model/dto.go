package model

import (
	"time"

	"github.com/google/uuid"
)

type TaskGroupData struct {
	Id        uuid.UUID
	Name      string
	DateAdded time.Time
}

type TaskData struct {
	Id          uuid.UUID
	Name        string
	Description string
	Content     []byte
	DateAdded   time.Time
}
