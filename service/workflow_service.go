package service

import "workflow-service/service/model/ioc"

type WorkflowService struct {
	WorkflowRepo ioc.IWorkflowRepository
}
