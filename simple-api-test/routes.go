package api

import (
	"github.com/gorilla/mux"
	"net/http"
)


type Route struct {
	Name    string
	Path    string
	Method  string
	Handler http.HandlerFunc
}

type Routes []Route

func (h *APIHandler) AddRoutesToGorillaMux(router *mux.Router) {
	for _, route := range h.GetRoutes() {
		router.
			Name(route.Name).
			Path(route.Path).
			Methods(route.Method).
			Handler(route.Handler)
	}
}

func (h *APIHandler) GetRoutes() Routes {
	return Routes{
		{
			"GETCineProjects",
			"/cineProjects",
			"GET",
			h.HandleGETCineProjects,
		},{
			"POSTCineProjects",
			"/cineProjects",
			"POST",
			h.HandlePOSTCineProjects,
		},{
			"DELETECineProjects_CineProjectId",
			"/cineProjects/{cineProjectId}",
			"DELETE",
			h.HandleDELETECineProjects_CineProjectId,
		},{
			"GETCineProjects_CineProjectId",
			"/cineProjects/{cineProjectId}",
			"GET",
			h.HandleGETCineProjects_CineProjectId,
		},{
			"PUTCineProjects_CineProjectId",
			"/cineProjects/{cineProjectId}",
			"PUT",
			h.HandlePUTCineProjects_CineProjectId,
		},
	}
}

