package service

import (
	"context"
	"strconv"
	"workflow-service/database"
	db_model "workflow-service/database/model"
	"workflow-service/service/model"
	"workflow-service/service/model/ioc"
	"workflow-service/transport/model/contracts"

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
func (srvc WorkflowService) CreateWorkflow(ctx context.Context, name string, appId uuid.UUID, requestData contracts.CreateWorkflowRequest) (id uuid.UUID, err error) {
	tx, err := srvc.TransactionManager.BeginTransaction(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	defer func() {
		err = srvc.TransactionManager.EndTransaction(tx, err)
	}()

	workflowData := dataFromWorkflowRequest(&requestData)
	id, err = srvc.WorkflowRepository.AddWorkflow(tx, name, appId, db_model.WorkflowSetting{
		IsFullPageRecognition: workflowData.IsFullPageRecognition,
		SkipImageEnhancement:  workflowData.SkipImageEnhancement,
		ExpectDifferentImages: workflowData.ExpectDifferentImages,
	})
	if err != nil {
		return uuid.Nil, err
	}
	workflowData.Id = id
	err = srvc.DaprService.PublishEvent(ctx, "new-workflow", workflowData)
	return
}

// Method for obtaining information about a specific workflow in the system.
func (srvc WorkflowService) GetWorkflowInfo(ctx context.Context, workflowId uuid.UUID) (db_model.WorkflowInfo, error) {
	return srvc.WorkflowRepository.GetWorkflow(workflowId)
}

// Method for obtaining a list of information about app's workflows.
func (srvc WorkflowService) GetWorkflowsInfo(ctx context.Context, appId uuid.UUID) ([]db_model.WorkflowInfo, error) {
	return srvc.WorkflowRepository.GetWorkflows(appId)
}

// A method for updating a specific workflow service.
func (srvc WorkflowService) UpdateWorkflow(ctx context.Context, name string, workflowId uuid.UUID, settings db_model.WorkflowSetting) error {
	tx, err := srvc.TransactionManager.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() { err = srvc.TransactionManager.EndTransaction(tx, err) }()

	err = srvc.WorkflowRepository.UpdateWorkflow(tx, workflowId, name, settings)
	if err != nil {
		return err
	}

	workflowData := model.WorkflowData{
		Id:                    workflowId,
		IsFullPageRecognition: settings.IsFullPageRecognition,
		SkipImageEnhancement:  settings.SkipImageEnhancement,
		ExpectDifferentImages: settings.ExpectDifferentImages,
	}
	err = srvc.DaprService.PublishEvent(ctx, "workflow_update", workflowData)
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
	if err == nil {
		err = srvc.DaprService.PublishEvent(ctx, "workflow_delete", workflowId)
	}
	return
}

func dataFromWorkflowRequest(request *contracts.CreateWorkflowRequest) *model.WorkflowData {
	isFullPage, _ := strconv.ParseBool(request.IsFullPageRecognition)
	expectDiffImages, _ := strconv.ParseBool(request.ExpectDifferentImages)
	skipEnhancement, _ := strconv.ParseBool(request.SkipImageEnhancement)
	return &model.WorkflowData{
		Id:                    uuid.New(),
		IsFullPageRecognition: isFullPage,
		ExpectDifferentImages: expectDiffImages,
		SkipImageEnhancement:  skipEnhancement,
	}
}
