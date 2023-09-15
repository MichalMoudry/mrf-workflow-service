package model

import (
	"time"

	"github.com/google/uuid"
)

// A structure encapsulating available info of a recognition app.
type ApplicationInfo struct {
	Id      uuid.UUID `db:"id"`
	Creator string    `db:"creator_id"`
	Name    string    `db:"app_name"`
	Added   time.Time `db:"date_added"`
}
