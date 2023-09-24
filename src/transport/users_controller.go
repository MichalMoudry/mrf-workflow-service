package transport

import (
	"io"
	"log"
	"net/http"
)

// A method for handling requests to delete all user's data.
func (handler *Handler) DeleteUsersData(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(body))
}
