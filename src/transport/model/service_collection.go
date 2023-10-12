package model

import (
	"workflow-service/database/repositories"
	"workflow-service/service"
	"workflow-service/transport/model/ioc"

	dapr "github.com/dapr/go-sdk/client"
)

// A structure representing a collection of public services in the service layer.
type ServiceCollection struct {
	AppService      ioc.IApplicationService
	WorkflowService ioc.IWorkflowService
	UserService     ioc.IUserService
}

// A constructor function for ServiceCollection structure.
func NewServiceCollection(daprClient dapr.Client) ServiceCollection {
	appRepository := &repositories.ApplicationRepository{}
	return ServiceCollection{
		AppService: service.NewAppService(
			appRepository,
		),
		WorkflowService: service.NewWorkflowService(
			repositories.WorkflowRepository{},
			daprClient,
		),
		UserService: service.NewUserService(
			appRepository,
		),
	}
}
