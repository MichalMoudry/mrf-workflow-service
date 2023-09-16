package ioc

import (
	"context"
	"workflow-service/database/model"

	"github.com/google/uuid"
)

// An interface for a recognition application service.
type IApplicationService interface {
	// A method for creating a new recognition app in the system.
	// This method returns app's id or error.
	CreateApp(ctx context.Context, name string) (uuid.UUID, error)

	// Method for retrieving information about a specific recognition app.
	GetAppInfo(ctx context.Context, appId uuid.UUID) (model.ApplicationInfo, error)

	// Method for retrieving information about user's applications.
	GetAppInfos(ctx context.Context) ([]model.ApplicationInfo, error)

	// A method for deleting an existing app from the system.
	DeleteApp(ctx context.Context, appId uuid.UUID) error

	// A method for updating a specific recognition app.
	UpdateApp(ctx context.Context, appId uuid.UUID, appName string) error
}
