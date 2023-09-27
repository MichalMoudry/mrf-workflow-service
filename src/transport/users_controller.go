package transport

import (
	"log"
	"net/http"
	"workflow-service/transport/model/contracts"
	"workflow-service/transport/util"
)

// A method for handling requests to delete all user's data.
func (handler *Handler) DeleteUsersData(w http.ResponseWriter, r *http.Request) {
	var eventData contracts.CloudEvent[string]
	err := util.UnmarshallRequest(r, &eventData)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	log.Println(eventData.Data) // eventData.Data = user id

	util.WriteResponse(w, http.StatusOK, nil)
}
