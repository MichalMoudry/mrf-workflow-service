package transport

import (
	"net/http"
	"time"
	"workflow-service/transport/errors"
	"workflow-service/transport/model/contracts"
	"workflow-service/transport/util"
)

// Method for handling requests for creating a new recognition app.
func (handler *Handler) CreateApp(w http.ResponseWriter, r *http.Request) {
	var requestData contracts.CreateAppRequest
	if err := util.UnmarshallRequest(r, &requestData); err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	id, err := handler.Services.AppService.CreateApp(
		r.Context(),
		requestData.Name,
	)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusCreated, contracts.CreatedAppResponse{
		Id:        id,
		Name:      requestData.Name,
		DateAdded: time.Now(),
	})
}

// Method for handling requests for obtaining info about a specific app.
func (handler *Handler) GetAppInfo(w http.ResponseWriter, r *http.Request) {
	appId, err := util.GetUuidFromUrl(r)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	appInfo, err := handler.Services.AppService.GetAppInfo(r.Context(), appId)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	util.WriteResponse(w, http.StatusOK, appInfo)
}

// Method for handling requests for obtaining a list of information about user's apps.
func (handler *Handler) GetUsersApps(w http.ResponseWriter, r *http.Request) {
	userId, ok := util.GetUserIdFromCtx(r.Context())
	if !ok {
		util.WriteErrResponse(w, http.StatusBadRequest, errors.ErrUidContextIssue)
		return
	}

	data, err := handler.Services.AppService.GetAppInfos(r.Context(), userId)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, data)
}

// Method for handling request for deleting a specific app.
func (handler *Handler) DeleteApp(w http.ResponseWriter, r *http.Request) {
	appId, err := util.GetUuidFromUrl(r)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	if err = handler.Services.AppService.DeleteApp(r.Context(), appId); err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, "App was deleted")
}

// Method for handling request for updating data of a specific app.
func (handler *Handler) UpdateApp(w http.ResponseWriter, r *http.Request) {
	appId, err := util.GetUuidFromUrl(r)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	var requestData contracts.UpdateAppRequest
	if err = util.UnmarshallRequest(r, &requestData); err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = handler.Services.AppService.UpdateApp(r.Context(), appId, requestData.Name); err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, "App was successfully updated.")
}
