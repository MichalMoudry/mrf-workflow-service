package ioc

import (
	"context"
	"workflow-service/database/model"

	"firebase.google.com/go/v4/auth"
	"github.com/google/uuid"
)

// An interface for a user service.
type IUserService interface {
	// Method for validating user's token.
	VerifyIdToken(ctx context.Context, token string) (*auth.Token, error)
}

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

// An interface for a recognition workflow service.
type IWorkflowService interface {
	// A method for creating a new workflow in the system.
	CreateWorkflow(ctx context.Context, name string, appId uuid.UUID, settings model.WorkflowSetting) (uuid.UUID, error)

	// Method for obtaining information about a specific workflow in the system.
	GetWorkflowInfo(ctx context.Context, workflowId uuid.UUID) (model.WorkflowInfo, error)

	// Method for obtaining a list of information about app's workflows.
	GetWorkflowsInfo(ctx context.Context, appId uuid.UUID) ([]model.WorkflowInfo, error)

	// Method for removing an existing service from the system.
	DeleteWorkflow(ctx context.Context, workflowId uuid.UUID) (err error)
}
