package service

import (
	"context"
	"workflow-service/database"
	"workflow-service/database/model"
	"workflow-service/service/model/ioc"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/google/uuid"
)

// A structure representing a service for working with the Workflow entity.
type WorkflowService struct {
	WorkflowRepository ioc.IWorkflowRepository
	TransactionManager ioc.ITransactionManager
	DaprService        ioc.IDaprService
}

// A constructor function for the Workflow structure.
func NewWorkflowService(workflowRepo ioc.IWorkflowRepository, dapr dapr.Client) WorkflowService {
	return WorkflowService{
		WorkflowRepository: workflowRepo,
		TransactionManager: database.TransactionManager{},
		DaprService:        NewDapr(dapr),
	}
}

// A method for creating a new workflow in the system.
func (srvc WorkflowService) CreateWorkflow(ctx context.Context, name string, appId uuid.UUID, settings model.WorkflowSetting) (id uuid.UUID, err error) {
	tx, err := srvc.TransactionManager.BeginTransaction(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	defer func() {
		err = srvc.TransactionManager.EndTransaction(tx, err)
	}()

	id, err = srvc.WorkflowRepository.AddWorkflow(tx, name, appId, settings)
	if err != nil {
		err = srvc.DaprService.PublishEvent(ctx, "new-workflow", id)
		if err != nil {
			return uuid.Nil, err
		}
	}
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

// A method for updating a specific workflow service.
func (srvc WorkflowService) UpdateWorkflow(ctx context.Context, name string, workflowId uuid.UUID, settings model.WorkflowSetting) error {
	tx, err := srvc.TransactionManager.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() { err = srvc.TransactionManager.EndTransaction(tx, err) }()

	err = srvc.WorkflowRepository.UpdateWorkflow(tx, workflowId, name, settings)
	if err != nil {
		return err
	}

	err = srvc.DaprService.PublishEvent(ctx, "workflow_update", settings)
	if err != nil {
		return err
	}
	return nil
}

// Method for removing an existing service from the system.
func (srvc WorkflowService) DeleteWorkflow(ctx context.Context, workflowId uuid.UUID) (err error) {
	tx, err := srvc.TransactionManager.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() { err = srvc.TransactionManager.EndTransaction(tx, err) }()
	err = srvc.WorkflowRepository.DeleteWorkflow(tx, workflowId)
	if err != nil {
		err = srvc.DaprService.PublishEvent(ctx, "workflow_delete", workflowId)
	}
	return
}
