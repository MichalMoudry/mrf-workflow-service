package ioc

import (
	"context"
	db_model "workflow-service/database/model"
	"workflow-service/service/model"
	"workflow-service/transport/model/contracts"

	"github.com/google/uuid"
)

// An interface for a recognition application service.
type IApplicationService interface {
	// A method for creating a new recognition app in the system.
	// This method returns app's id or error.
	CreateApp(ctx context.Context, userId string, name string) (uuid.UUID, error)

	// Method for retrieving information about a specific recognition app.
	GetAppInfo(ctx context.Context, appId uuid.UUID) (db_model.ApplicationInfo, error)

	// Method for retrieving information about user's applications.
	GetAppInfos(ctx context.Context, userId string) ([]db_model.ApplicationInfo, error)

	// A method for deleting an existing app from the system.
	DeleteApp(ctx context.Context, appId uuid.UUID) error

	// A method for updating a specific recognition app.
	UpdateApp(ctx context.Context, appId uuid.UUID, appName string) error
}

// An interface for a recognition workflow service.
type IWorkflowService interface {
	// A method for creating a new workflow in the system.
	CreateWorkflow(ctx context.Context, name string, appId uuid.UUID, requestData contracts.CreateWorkflowRequest) (uuid.UUID, error)

	// Method for obtaining information about a specific workflow in the system.
	GetWorkflowInfo(ctx context.Context, workflowId uuid.UUID) (db_model.WorkflowInfo, error)

	// Method for obtaining a list of information about app's workflows.
	GetWorkflowsInfo(ctx context.Context, appId uuid.UUID) ([]db_model.WorkflowInfo, error)

	// A method for updating a specific workflow service.
	UpdateWorkflow(ctx context.Context, name string, workflowId uuid.UUID, settings db_model.WorkflowSetting) error

	// Method for removing an existing service from the system.
	DeleteWorkflow(ctx context.Context, workflowId uuid.UUID) (err error)
}

// An interface for a user service.
type IUserService interface {
	// Method for deleting all user's data in the system.
	DeleteUsersData(ctx context.Context, userId string) error
}

// An interface for a tasks service.
type ITasksService interface {
	// A method for obtaining a list of task groups connected to a specific workflow.
	GetTaskGroups(ctx context.Context, workflowId uuid.UUID) ([]model.TaskGroupData, error)

	// A method for obtaining a list of tasks for a specific task group.
	GetTasks(ctx context.Context, groupId uuid.UUID) ([]model.TaskData, error)
}
