package repositories

import (
	"workflow-service/database"
	"workflow-service/database/model"
	"workflow-service/database/query"

	"github.com/google/uuid"
)

type WorkflowRepository struct{}

// A method for adding a new workflow to the database.
func (WorkflowRepository) AddWorkflow(name string, appId uuid.UUID, settings model.WorkflowSetting) (uuid.UUID, error) {
	ctx, err := database.GetDbContext()
	if err != nil {
		return uuid.Nil, err
	}

	rows, err := ctx.NamedQuery(query.CreateWorkflow, model.NewWorkflow(name, appId, settings))
	if err != nil {
		return uuid.Nil, nil
	}
	var returnedId string
	for rows.Next() {
		err = rows.Scan(&returnedId)
	}
	if err != nil {
		return uuid.Nil, nil
	}

	return uuid.Parse(returnedId)
}

// Method for retrieving info about a workflow from the database.
func (WorkflowRepository) GetWorkflow(workflowId uuid.UUID) (model.WorkflowInfo, error) {
	ctx, err := database.GetDbContext()
	if err != nil {
		return model.WorkflowInfo{}, err
	}

	var data model.WorkflowInfo
	if err = ctx.Get(&data, query.GetWorkflow, workflowId); err != nil {
		return model.WorkflowInfo{}, nil
	}
	return data, nil
}

// A method for retrieving a list of workflows for a specific recognition app.
func (WorkflowRepository) GetWorkflows(appId uuid.UUID) ([]model.WorkflowInfo, error) {
	ctx, err := database.GetDbContext()
	if err != nil {
		return nil, err
	}

	var workflows []model.WorkflowInfo
	if err = ctx.Select(&workflows, query.GetWorkflows, appId); err != nil {
		return nil, err
	}
	return workflows, nil
}

// Method for deleting a workflow from the database.
func (WorkflowRepository) DeleteWorkflow(workflowId uuid.UUID) error {
	ctx, err := database.GetDbContext()
	if err != nil {
		return err
	}

	if _, err = ctx.Exec(query.DeleteWorkflow, workflowId); err != nil {
		return err
	}
	return nil
}
