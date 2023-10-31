package ioc

import (
	"workflow-service/database/model"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
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

	// Method for deleting all user's recognition applications.
	DeleteUsersApps(tx *sqlx.Tx, userId string) error

	// Method for updating app's name.
	UpdateApplication(appId uuid.UUID, app_name string) error
}

// An interface for a repository connected to Workflow structure.
type IWorkflowRepository interface {
	// A method for adding a new workflow to the database.
	AddWorkflow(tx *sqlx.Tx, name string, appId uuid.UUID, settings model.WorkflowSetting) (uuid.UUID, error)

	// Method for retrieving info about a workflow from the database.
	GetWorkflow(workflowId uuid.UUID) (model.WorkflowInfo, error)

	// A method for retrieving a list of workflows for a specific recognition app.
	GetWorkflows(appId uuid.UUID) ([]model.WorkflowInfo, error)

	// A method for updating a specific recognition workflow in the database.
	UpdateWorkflow(tx *sqlx.Tx, id uuid.UUID, name string, settings model.WorkflowSetting) error

	// Method for deleting a workflow from the database.
	DeleteWorkflow(tx *sqlx.Tx, workflowId uuid.UUID) error
}

// An interface for a repository connected to processing tasks.
type ITasksRepository interface {
	// A method for obtaining a list of task groups.
	GetTaskGroups(workflowId uuid.UUID) ([]model.TaskGroup, error)

	// A method for obtaining a list of tasks.
	GetTasks(groupId uuid.UUID) ([]model.Task, error)
}
