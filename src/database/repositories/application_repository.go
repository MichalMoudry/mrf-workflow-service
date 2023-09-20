package repositories

import (
	"time"
	"workflow-service/database"
	"workflow-service/database/model"
	"workflow-service/database/query"

	"github.com/google/uuid"
)

type ApplicationRepository struct{}

// A method for creating a new application in the database.
func (ApplicationRepository) AddApplication(name, creatorId string) (uuid.UUID, error) {
	ctx, err := database.GetDbContext()
	if err != nil {
		return uuid.Nil, err
	}

	rows, err := ctx.NamedQuery(query.CreateApp, model.NewApplication(name, creatorId))
	if err != nil {
		return uuid.Nil, err
	}
	var returnedId string
	for rows.Next() {
		err = rows.Scan(&returnedId)
	}
	if err != nil {
		return uuid.Nil, err
	}

	return uuid.Parse(returnedId)
}

// A method for retrieving basic info about a specific app.
func (ApplicationRepository) GetApp(appId uuid.UUID) (model.ApplicationInfo, error) {
	ctx, err := database.GetDbContext()
	if err != nil {
		return model.ApplicationInfo{}, err
	}

	var data model.ApplicationInfo
	if err = ctx.Get(&data, query.GetApp, appId); err != nil {
		return model.ApplicationInfo{}, nil
	}
	return data, nil
}

// A method for retrieving basic info about user's apps from the database.
func (ApplicationRepository) GetUsersApps(userId string) ([]model.ApplicationInfo, error) {
	ctx, err := database.GetDbContext()
	if err != nil {
		return nil, err
	}

	var apps []model.ApplicationInfo
	if err = ctx.Select(&apps, query.GetApps, userId); err != nil {
		return nil, err
	}
	return apps, nil
}

// A method for deleting an existing app from the database.
func (ApplicationRepository) DeleteApplication(appId uuid.UUID) error {
	ctx, err := database.GetDbContext()
	if err != nil {
		return err
	}

	if _, err = ctx.Exec(query.DeleteApp, appId); err != nil {
		return err
	}
	return nil
}

// Method for updating app's name.
func (ApplicationRepository) UpdateApplication(appId uuid.UUID, appName string, updateTime time.Time) error {
	ctx, err := database.GetDbContext()
	if err != nil {
		return err
	}

	if _, err = ctx.Exec(query.UpdateApp, appId, appName, updateTime); err != nil {
		return err
	}
	return nil
}
