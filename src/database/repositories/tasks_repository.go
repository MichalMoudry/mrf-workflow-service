package repositories

import (
	"workflow-service/database"
	"workflow-service/database/model"
	"workflow-service/database/query"

	"github.com/google/uuid"
)

type TasksRepository struct{}

// A method for obtaining a list of task groups.
func (TasksRepository) GetTaskGroups(workflowId uuid.UUID) ([]model.TaskGroup, error) {
	ctx, err := database.GetDbContext()
	if err != nil {
		return nil, err
	}
	var taskGroups []model.TaskGroup
	if err = ctx.Select(&taskGroups, query.GetTaskGroups, workflowId); err != nil {
		return nil, err
	}

	return taskGroups, nil
}

// A method for obtaining a list of tasks.
func (TasksRepository) GetTasks(groupId uuid.UUID) ([]model.Task, error) {
	ctx, err := database.GetDbContext()
	if err != nil {
		return nil, err
	}
	var tasks []model.Task
	if err = ctx.Select(&tasks, query.GetTasks, groupId); err != nil {
		return nil, err
	}

	return tasks, nil
}
