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
