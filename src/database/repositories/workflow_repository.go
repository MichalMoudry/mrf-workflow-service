package repositories

import (
	"time"
	"workflow-service/database"
	"workflow-service/database/model"
	"workflow-service/database/query"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type WorkflowRepository struct{}

// A method for adding a new workflow to the database.
func (WorkflowRepository) AddWorkflow(tx *sqlx.Tx, name string, appId uuid.UUID, settings model.WorkflowSetting) (uuid.UUID, error) {
	rows, err := tx.NamedQuery(query.CreateWorkflow, model.NewWorkflow(name, appId, settings))
	if err != nil {
		return uuid.Nil, err
	}
	var returnedId string
	for rows.Next() {
		err = rows.Scan(&returnedId)
	}
	if err != nil {
		return uuid.Nil, err
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
		return model.WorkflowInfo{}, err
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

// A method for updating a specific recognition workflow in the database.
func (WorkflowRepository) UpdateWorkflow(tx *sqlx.Tx, id uuid.UUID, name string, settings model.WorkflowSetting) error {
	_, err := tx.Exec(
		query.UpdateWorkflow,
		id,
		name,
		settings.IsFullPageRecognition,
		settings.SkipImageEnhancement,
		settings.ExpectDifferentImages,
		uuid.New(),
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}

// Method for deleting a workflow from the database.
func (WorkflowRepository) DeleteWorkflow(tx *sqlx.Tx, workflowId uuid.UUID) error {
	if _, err := tx.Exec(query.DeleteWorkflow, workflowId); err != nil {
		return err
	}
	return nil
}
