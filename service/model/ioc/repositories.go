package ioc

import (
	"time"
	"workflow-service/database/model"

	"github.com/google/uuid"
)

// An interface for a repository connected to Application structure.
type IApplicationRepository interface {
	// A method for creating a new application in the database.
	AddApplication(name, creatorId string) (uuid.UUID, error)

	// A method for retrieving basic info about a specific app.
	GetApp(appId uuid.UUID) (model.ApplicationInfo, error)

	// A method for retrieving basic info about user's apps from the database.
	GetUsersApps(userId string) ([]model.ApplicationInfo, error)

	// A method for deleting an existing app from the database.
	DeleteApplication(appId uuid.UUID) error

	// Method for updating app's name.
	UpdateApplication(appId uuid.UUID, app_name string, updateTime time.Time) error
}
