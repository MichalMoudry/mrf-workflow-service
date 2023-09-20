package service

import (
	"context"
	"workflow-service/database"
	"workflow-service/database/model"
	"workflow-service/service/model/ioc"

	"github.com/google/uuid"
)

// A structure representing a service for working with the Workflow entity.
type WorkflowService struct {
	WorkflowRepository ioc.IWorkflowRepository
}

// A constructor function for the Workflow structure.
func NewWorkflowService(workflowRepo ioc.IWorkflowRepository) WorkflowService {
	return WorkflowService{
		WorkflowRepository: workflowRepo,
	}
}

// A method for creating a new workflow in the system.
func (srvc WorkflowService) CreateWorkflow(ctx context.Context, name string, appId uuid.UUID, settings model.WorkflowSetting) (id uuid.UUID, err error) {
	tx, err := database.BeginTransaction(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	defer func() {
		err = database.EndTransaction(tx, err)
	}()

	id, err = srvc.WorkflowRepository.AddWorkflow(name, appId, settings)
	return
}

// Method for obtaining information about a specific workflow in the system.
func (srvc WorkflowService) GetWorkflowInfo(ctx context.Context, workflowId uuid.UUID) (model.WorkflowInfo, error) {
	return srvc.WorkflowRepository.GetWorkflow(workflowId)
}

// Method for obtaining a list of information about app's workflows.
func (srvc WorkflowService) GetWorkflowsInfo(ctx context.Context, appId uuid.UUID) ([]model.WorkflowInfo, error) {
	return srvc.WorkflowRepository.GetWorkflows(appId)
}

// Method for removing an existing service from the system.
func (srvc WorkflowService) DeleteWorkflow(ctx context.Context, workflowId uuid.UUID) (err error) {
	tx, err := database.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() { err = database.EndTransaction(tx, err) }()
	err = srvc.WorkflowRepository.DeleteWorkflow(workflowId)
	return
}