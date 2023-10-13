package transport

import (
	"net/http"
	"workflow-service/transport/model"

	srvc_middleware "workflow-service/transport/middleware"

	"firebase.google.com/go/v4/auth"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	Port     int
	Mux      *chi.Mux
	Services model.ServiceCollection
}

// Initializer function for HTTP handler.
func Initalize(port int, services model.ServiceCollection, auth *auth.Client) *Handler {
	handler := &Handler{
		Port:     port,
		Mux:      chi.NewRouter(),
		Services: services,
	}
	handler.Mux.Use(middleware.Logger)

	//Protected routes
	handler.Mux.Group(func(r chi.Router) {
		r.Use(srvc_middleware.Authenticate(auth))

		r.Route("/apps", func(r chi.Router) {
			r.Post("/", handler.CreateApp)
			r.Get("/", handler.GetUsersApps)
			r.Route("/{uuid}", func(r chi.Router) {
				r.Get("/", handler.GetAppInfo)
				r.Delete("/", handler.DeleteApp)
				r.Patch("/", handler.UpdateApp)
			})
		})

		r.Route("/workflows", func(r chi.Router) {
			r.Post("/", handler.CreateWorkflow)
			r.Get("/app/{uuid}", handler.GetWorkflowsInfo)
			r.Route("/{uuid}", func(r chi.Router) {
				r.Get("/", handler.GetWorkflowInfo)
				r.Patch("/", handler.UpdateWorkflow)
				r.Delete("/", handler.DeleteWorkflow)
			})
		})

		r.Route("/taskgroups", func(r chi.Router) {
			r.Route("/{uuid}", func(r chi.Router) {
				r.Delete("/", handler.DeleteTaskGroup)
				r.Patch("/", handler.PatchTaskGroup)
			})
		})

		r.Route("/tasks", func(r chi.Router) {

		})

		r.Route("/users", func(r chi.Router) {
			r.Post("/delete", handler.DeleteUsersData)
		})
	})

	// Public routes
	handler.Mux.Get("/health", health)

	// Dapr routes
	handler.Mux.Get("/dapr/subscribe", ConfigureSubscribeHandler)
	return handler
}

// Controller endpoint function for handling requests on /health.
func health(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNoContent)
}
