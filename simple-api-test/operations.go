package api

import (
	"context"
	"github.com/rs/zerolog"
	"net/http"
	"os"
	"time"
)


// APIHandler is a type to give the api functions below access to a common logger
// any any other shared objects
type APIHandler struct {
	// Zerolog was chosen as the default logger, but you can replace it with any logger of your choice
	logger zerolog.Logger

	// Note: if you need to pass in a client for your database, this would be a good place to include it
}

func NewAPIHandler() *APIHandler {
	output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	logger := zerolog.New(output).With().Timestamp().Logger()
	return &APIHandler{logger: logger}
}

func (h *APIHandler) WithLogger(logger zerolog.Logger) *APIHandler {
	h.logger = logger
	return h
}

// Deletes a specific cine project
func (h *APIHandler) DELETECineProjects_CineProjectId(ctx context.Context, cineProjectId int64) (Response, error) {
	// TODO: implement the DELETECineProjects_CineProjectId function to return the following responses

	// return NewResponse(204, {}, "", responseHeaders), nil

	// return NewResponse(404, {}, "", responseHeaders), nil

	return NewResponse(http.StatusNotImplemented, ErrorMsg{"DELETECineProjects_CineProjectId operation has not been implemented yet"}, "application/json", nil), nil
}

// Returns a list of cine projects
func (h *APIHandler) GETCineProjects(ctx context.Context) (Response, error) {
	// TODO: implement the GETCineProjects function to return the following responses

	// return NewResponse(200, []CineProject, "application/json", responseHeaders), nil

	return NewResponse(http.StatusNotImplemented, ErrorMsg{"GETCineProjects operation has not been implemented yet"}, "application/json", nil), nil
}

// Returns a specific cine project
func (h *APIHandler) GETCineProjects_CineProjectId(ctx context.Context, cineProjectId int64) (Response, error) {
	// TODO: implement the GETCineProjects_CineProjectId function to return the following responses

	// return NewResponse(200, CineProject{}, "application/json", responseHeaders), nil

	// return NewResponse(404, {}, "", responseHeaders), nil

	return NewResponse(http.StatusNotImplemented, ErrorMsg{"GETCineProjects_CineProjectId operation has not been implemented yet"}, "application/json", nil), nil
}

// Creates a new cine project
func (h *APIHandler) POSTCineProjects(ctx context.Context, reqBody CineProject) (Response, error) {
	// TODO: implement the POSTCineProjects function to return the following responses

	// return NewResponse(201, CineProject{}, "application/json", responseHeaders), nil

	// return NewResponse(400, {}, "", responseHeaders), nil

	return NewResponse(http.StatusNotImplemented, ErrorMsg{"POSTCineProjects operation has not been implemented yet"}, "application/json", nil), nil
}

// Updates a specific cine project
func (h *APIHandler) PUTCineProjects_CineProjectId(ctx context.Context, cineProjectId int64, reqBody CineProject) (Response, error) {
	// TODO: implement the PUTCineProjects_CineProjectId function to return the following responses

	// return NewResponse(200, CineProject{}, "application/json", responseHeaders), nil

	// return NewResponse(400, {}, "", responseHeaders), nil

	// return NewResponse(404, {}, "", responseHeaders), nil

	return NewResponse(http.StatusNotImplemented, ErrorMsg{"PUTCineProjects_CineProjectId operation has not been implemented yet"}, "application/json", nil), nil
}

