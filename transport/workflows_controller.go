package transport

import (
	"net/http"
	"workflow-service/database/model"
	"workflow-service/transport/model/contracts"
	"workflow-service/transport/util"
)

// Method for handling requests for creating a new recognition workflow.
func (handler *Handler) CreateWorkflow(w http.ResponseWriter, r *http.Request) {
	var requestData contracts.CreateWorkflowRequest
	if err := util.UnmarshallRequest(r, &requestData); err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	workflowId, err := handler.Services.WorkflowService.CreateWorkflow(
		r.Context(),
		requestData.Name,
		util.ParseStringAsUuid(requestData.AppId),
		model.WorkflowSetting{
			IsFullPageRecognition: requestData.IsFullPageRecognition,
			SkipImageEnhancement:  requestData.SkipImageEnhancement,
			ExpectDifferentImages: requestData.ExpectDifferentImages,
		},
	)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusCreated, workflowId)
}

// Method for handling requests for obtaining information about app's workflows.
func (handler *Handler) GetWorkflowsInfo(w http.ResponseWriter, r *http.Request) {
	appId, err := util.GetUuidFromUrl(r)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	data, err := handler.Services.WorkflowService.GetWorkflowsInfo(
		r.Context(),
		appId,
	)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	util.WriteResponse(w, http.StatusOK, data)
}

// Method for handling requests for obtaining information about a single workflow.
func (handler *Handler) GetWorkflowInfo(w http.ResponseWriter, r *http.Request) {
	workflowId, err := util.GetUuidFromUrl(r)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	data, err := handler.Services.WorkflowService.GetWorkflowInfo(r.Context(), workflowId)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	util.WriteResponse(w, http.StatusOK, data)
}

// Method for handling requests to delete a specific recognition workflow.
func (handler *Handler) DeleteWorkflow(w http.ResponseWriter, r *http.Request) {
	workflowId, err := util.GetUuidFromUrl(r)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = handler.Services.WorkflowService.DeleteWorkflow(r.Context(), workflowId); err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	util.WriteResponse(w, http.StatusOK, nil)
}
