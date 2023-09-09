package transport

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	Port int
	Mux  *chi.Mux
}

// Initializer function for HTTP handler.
func Initalize(port int) *Handler {
	handler := &Handler{
		Port: port,
		Mux:  chi.NewRouter(),
	}
	handler.Mux.Use(middleware.Logger)

	// Public routes
	handler.Mux.Get("/healthz", health)
	return handler
}

// Controller endpoint function for handling requests on /healthz.
func health(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNoContent)
}