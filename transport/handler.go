package transport

import (
	"net/http"
	"workflow-service/transport/model"
	"workflow-service/transport/util"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/docgen"
)

type Handler struct {
	Port     int
	Mux      *chi.Mux
	Services model.ServiceCollection
}

// Initializer function for HTTP handler.
func Initalize(port int, services model.ServiceCollection) *Handler {
	handler := &Handler{
		Port:     port,
		Mux:      chi.NewRouter(),
		Services: services,
	}
	handler.Mux.Use(middleware.Logger)

	//Protected routes
	handler.Mux.Group(func(r chi.Router) {
		r.Route("/apps", func(r chi.Router) {
			r.Post("/", handler.CreateApp)
			r.Get("/", handler.GetUsersApps)
			r.Route("/{appid}", func(r chi.Router) {
				r.Get("/", handler.GetAppInfo)
				r.Delete("/", handler.DeleteApp)
				r.Patch("/", handler.UpdateApp)
			})
		})

		r.Route("/workflows", func(r chi.Router) {
			r.Post("/", handler.CreateWorkflow)
			r.Get("/app/{appid}", handler.GetWorkflowsInfo)
			r.Route("/{workflowid}", func(r chi.Router) {
				r.Get("/", handler.GetWorkflowInfo)
				r.Delete("/", handler.DeleteWorkflow)
			})
		})

		r.Route("/templates", func(r chi.Router) {

		})
	})

	// Public routes
	handler.Mux.Get("/health", health)
	handler.Mux.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
		util.WriteResponse(w, http.StatusOK, docgen.JSONRoutesDoc(handler.Mux))
	})
	docgen.JSONRoutesDoc(handler.Mux)
	return handler
}

// Controller endpoint function for handling requests on /health.
func health(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNoContent)
}
