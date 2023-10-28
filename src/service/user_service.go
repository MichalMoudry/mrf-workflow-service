package service

import (
	"context"
	"workflow-service/database"
	"workflow-service/service/model/ioc"
)

type UserService struct {
	AppRepository      ioc.IApplicationRepository
	TransactionManager ioc.ITransactionManager
}

// A constructor function for the UserService structure.
func NewUserService(appRepo ioc.IApplicationRepository) *UserService {
	return &UserService{
		AppRepository:      appRepo,
		TransactionManager: database.TransactionManager{},
	}
}

// Method for deleting all user's data in the system.
func (srvc UserService) DeleteUsersData(ctx context.Context, userId string) (err error) {
	tx, err := srvc.TransactionManager.BeginTransaction(ctx)
	if err != nil {
		return
	}
	defer func() { err = srvc.TransactionManager.EndTransaction(tx, err) }()

	err = srvc.AppRepository.DeleteUsersApps(tx, userId)
	return
}
