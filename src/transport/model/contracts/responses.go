package contracts

import (
	"time"

	"github.com/google/uuid"
)

// A structure representing a structured response after creating a recognition app.
type CreatedAppResponse struct {
	Id        uuid.UUID
	Name      string
	DateAdded time.Time
}
