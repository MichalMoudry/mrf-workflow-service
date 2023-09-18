package model

import (
	"workflow-service/database/repositories"
	"workflow-service/service"
	"workflow-service/transport/model/ioc"

	"firebase.google.com/go/v4/auth"
)

// A structure representing a collection of public services in the service layer.
type ServiceCollection struct {
	UserService     ioc.IUserService
	AppService      ioc.IApplicationService
	WorkflowService ioc.IWorkflowService
	TemplateService ioc.ITemplateService
}

// A constructor function for ServiceCollection structure.
func NewServiceCollection(authClient *auth.Client) ServiceCollection {
	return ServiceCollection{
		AppService: service.NewAppService(
			repositories.ApplicationRepository{},
		),
		WorkflowService: service.NewWorkflowService(
			repositories.WorkflowRepository{},
		),
		UserService: service.NewUserService(authClient),
	}
}
