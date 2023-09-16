package transport

import (
	"net/http"
	"time"
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
