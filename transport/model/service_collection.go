package model

import (
	"workflow-service/database/repositories"
	"workflow-service/service"
	"workflow-service/transport/model/ioc"
)

// A structure representing a collection of public services in the service layer.
type ServiceCollection struct {
	AppService ioc.IApplicationService
}

// A constructor function for ServiceCollection structure.
func NewServiceCollection() ServiceCollection {
	return ServiceCollection{
		AppService: service.NewAppService(
			repositories.ApplicationRepository{},
		),
	}
}
