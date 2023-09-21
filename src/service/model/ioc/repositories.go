package ioc

import (
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
	UpdateApplication(appId uuid.UUID, app_name string) error
}

// An interface for a repository connected to Workflow structure.
type IWorkflowRepository interface {
	// A method for adding a new workflow to the database.
	AddWorkflow(name string, appId uuid.UUID, settings model.WorkflowSetting) (uuid.UUID, error)

	// Method for retrieving info about a workflow from the database.
	GetWorkflow(workflowId uuid.UUID) (model.WorkflowInfo, error)

	// A method for retrieving a list of workflows for a specific recognition app.
	GetWorkflows(appId uuid.UUID) ([]model.WorkflowInfo, error)

	// Method for deleting a workflow from the database.
	DeleteWorkflow(workflowId uuid.UUID) error
}

// An interface for a repository connected to document templates.
type ITemplateRepository interface {
	// A method for adding a new document template to the database.
	AddTemplate(template *model.DocumentTemplate) (uuid.UUID, error)

	// A method for deleting a specific template from the database.
	DeleteTemplate(templateId uuid.UUID) error
}

// An interface for a repository connected to document template fields.
type IFieldRepository interface {
	// Method for adding document template fields in bulk.
	AddFields(fields []*model.TemplateField) error

	// Method for bulk deleting document template fields.
	DeleteFields(fieldIds []uuid.UUID) error
}
