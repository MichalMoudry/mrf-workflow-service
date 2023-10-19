package service

import (
	"context"
	"workflow-service/service/model"
	"workflow-service/service/model/ioc"

	"github.com/google/uuid"
)

type TasksService struct {
	TasksRepository ioc.ITasksRepository
}

// A constructor function for the TasksService structure.
func NewTasksService(tasksRepo ioc.ITasksRepository) *TasksService {
	return &TasksService{
		TasksRepository: tasksRepo,
	}
}

// A method for obtaining a list of task groups connected to a specific workflow.
func (srvc TasksService) GetTaskGroups(ctx context.Context, workflowId uuid.UUID) ([]model.TaskGroupData, error) {
	data, err := srvc.TasksRepository.GetTaskGroups(workflowId)
	if err != nil {
		return nil, err
	}
	groupNumber := len(data)
	result := make([]model.TaskGroupData, groupNumber)
	if groupNumber == 0 {
		return result, nil
	}

	for i, v := range data {
		result[i] = model.TaskGroupData{
			Id:        v.Id,
			Name:      v.Name,
			DateAdded: v.DateAdded,
		}
	}
	return result, nil
}

// A method for obtaining a list of tasks for a specific task group.
func (srvc TasksService) GetTasks(ctx context.Context, groupId uuid.UUID) ([]model.TaskData, error) {
	data, err := srvc.TasksRepository.GetTasks(groupId)
	if err != nil {
		return nil, err
	}
	groupNumber := len(data)
	result := make([]model.TaskData, groupNumber)
	if groupNumber == 0 {
		return result, nil
	}

	for i, v := range data {
		result[i] = model.TaskData{
			Id:          v.Id,
			Name:        v.Name,
			Description: v.Description,
			Content:     v.Content,
			DateAdded:   v.DateAdded,
		}
	}
	return result, nil
}
