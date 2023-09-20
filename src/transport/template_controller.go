package transport

import (
	"io"
	"net/http"
	"workflow-service/transport/model/contracts"
	"workflow-service/transport/util"
)

// Method for handling requests for creating a new document template.
func (handler *Handler) CreateTemplate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	file, _, err := r.FormFile("template")
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	defer file.Close()

	var requestData contracts.CreateTemplateRequest
	if err = util.UnmarshallRequest(r, &requestData); err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	_, err = io.ReadAll(file)
	if err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	util.WriteResponse(w, http.StatusCreated, nil)
}

// Method for handling requests for deleting a specific document template.
func (handler *Handler) DeleteTemplate(w http.ResponseWriter, r *http.Request) {
	templateId, err := util.GetUuidFromUrl(r)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = handler.Services.TemplateService.DeleteTemplate(r.Context(), templateId); err != nil {
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	util.WriteResponse(w, http.StatusOK, nil)
}

// Method for handling requests for image of a specific document template.
func (handler *Handler) UpdateTemplateImage(w http.ResponseWriter, r *http.Request) {
	_, err := util.GetUuidFromUrl(r)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, nil)
}
