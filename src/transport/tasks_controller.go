package transport

import (
	"net/http"
	"workflow-service/transport/model/contracts"
	"workflow-service/transport/util"
)

// A method for handling requests for creating a new task group.
func (handler *Handler) CreateTaskGroup(w http.ResponseWriter, r *http.Request) {
	var requestData contracts.CreateTaskGroupRequest
	err := util.UnmarshallRequest(r, &requestData)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	util.WriteResponse(w, http.StatusCreated, nil)
}

// A method for handling request for obtaining task groups for a specific workflow.
func (handler *Handler) GetTaskGroups(w http.ResponseWriter, r *http.Request) {
	var requestData contracts.GetTaskGroupsRequest
	err := util.UnmarshallRequest(r, &requestData)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, nil)
}

// A method for handling request for patching a specific task group
func (handler *Handler) PatchTaskGroup(w http.ResponseWriter, r *http.Request) {
	var requestData contracts.PatchTaskGroupRequest
	err := util.UnmarshallRequest(r, &requestData)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, nil)
}

// A method for handling request for deleting a specific task group.
func (handler *Handler) DeleteTaskGroup(w http.ResponseWriter, r *http.Request) {
	_, err := util.GetUuidFromUrl(r)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	util.WriteErrResponse(w, http.StatusOK, nil)
}

// A method for handling request for creating a new task.
func (handler *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var requestData contracts.CreateTaskRequest
	err := util.UnmarshallRequest(r, &requestData)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	util.WriteResponse(w, http.StatusCreated, nil)
}
