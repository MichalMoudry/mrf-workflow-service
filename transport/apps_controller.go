package transport

import (
	"net/http"
	"time"
	"workflow-service/transport/model/contracts"
	"workflow-service/transport/util"
)

// Method for handling requests for creating a new recognition app.
func (handler *Handler) CreateApp(writer http.ResponseWriter, request *http.Request) {
	var requestData contracts.CreateAppRequestData
	if err := util.UnmarshallRequest(request, &requestData); err != nil {
		util.WriteErrResponse(writer, http.StatusBadRequest, err)
		return
	}

	id, err := handler.Services.AppService.CreateApp(
		request.Context(),
		requestData.Name,
	)
	if err != nil {
		util.WriteErrResponse(writer, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(writer, http.StatusCreated, contracts.CreatedAppResponse{
		Id:        id,
		Name:      requestData.Name,
		DateAdded: time.Now(),
	})
}
