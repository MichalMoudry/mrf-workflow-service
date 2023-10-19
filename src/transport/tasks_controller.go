package transport

import (
	"net/http"
	"workflow-service/transport/model/contracts"
	"workflow-service/transport/util"

	"github.com/google/uuid"
)

// A method for handling request for obtaining task groups for a specific workflow.
func (handler *Handler) GetTaskGroups(w http.ResponseWriter, r *http.Request) {
	var requestData contracts.GetTaskGroupsRequest
	err := util.UnmarshallRequest(r, &requestData)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	data, err := handler.Services.TasksService.GetTaskGroups(
		r.Context(),
		uuid.MustParse(requestData.WorkflowId),
	)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	util.WriteResponse(w, http.StatusOK, data)
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

	util.WriteResponse(w, http.StatusOK, nil)
}

// A method for handling requests for obtaining tasks for a specific task group.
func (handler *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	groupId, err := util.GetUuidFromUrl(r)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	data, err := handler.Services.TasksService.GetTasks(r.Context(), groupId)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	util.WriteResponse(w, http.StatusOK, data)
}
