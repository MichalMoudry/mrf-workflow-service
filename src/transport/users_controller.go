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
	log.Println("User delete started...")
	if err := util.UnmarshallRequest(r, &eventData); err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	log.Printf("Obtained an user-delete event for %v\n", eventData.Data)
	err := handler.Services.UserService.DeleteUsersData(r.Context(), eventData.Data)
	if err != nil {
		log.Println("User delete failed")
		util.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}

	util.WriteResponse(w, http.StatusOK, nil)
}
