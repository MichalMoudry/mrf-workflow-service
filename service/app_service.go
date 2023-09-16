package service

import (
	"context"
	"time"
	"workflow-service/database"
	"workflow-service/database/model"
	"workflow-service/service/model/ioc"
	"workflow-service/service/util"

	"github.com/google/uuid"
)

// A structure representing a service for working with Application entity.
type ApplicationService struct {
	AppRepository ioc.IApplicationRepository
}

// A constructor function for ApplicationService structure.
func NewAppService(appRepo ioc.IApplicationRepository) *ApplicationService {
	return &ApplicationService{
		AppRepository: appRepo,
	}
}

// A method for creating a new recognition app in the system.
// This method returns app's id or error.
func (srvc ApplicationService) CreateApp(ctx context.Context, name string) (uuid.UUID, error) {
	userId, err := util.GetUserIdFromContext(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	tx, err := database.BeginTransaction(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	defer func() {
		err = database.EndTransaction(tx, err)
	}()

	appId, err := srvc.AppRepository.AddApplication(name, userId)
	if err != nil {
		return uuid.Nil, err
	}
	return appId, nil
}

// Method for retrieving information about a specific recognition app.
func (srvc ApplicationService) GetAppInfo(ctx context.Context, appId uuid.UUID) (model.ApplicationInfo, error) {
	_, err := util.GetUserIdFromContext(ctx)
	if err != nil {
		return model.ApplicationInfo{}, err
	}

	appData, err := srvc.AppRepository.GetApplication(appId)
	if err != nil {
		return model.ApplicationInfo{}, err
	}
	return appData, nil
}

// A method for deleting an existing app from the system.
func (srvc ApplicationService) DeleteApp(ctx context.Context, appId uuid.UUID) error {
	_, err := util.GetUserIdFromContext(ctx)
	if err != nil {
		return err
	}

	if err = srvc.AppRepository.DeleteApplication(appId); err != nil {
		return err
	}
	return nil
}

// A method for updating a specific recognition app.
func (srvc ApplicationService) UpdateApp(ctx context.Context, appId uuid.UUID, appName string) error {
	_, err := util.GetUserIdFromContext(ctx)
	if err != nil {
		return err
	}

	if err = srvc.AppRepository.UpdateApplication(appId, appName, time.Now()); err != nil {
		return err
	}
	return nil
}
