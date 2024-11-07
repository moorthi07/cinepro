package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// HandleDELETECineProjects_CineProjectId handles parsing input to pass to the DELETECineProjects_CineProjectId operation and sends responses back to the client
func (h *APIHandler) HandleDELETECineProjects_CineProjectId(w http.ResponseWriter, r *http.Request) {
	var err error
	pathParams := mux.Vars(r)

	var cineProjectId int64
	cineProjectIdPathParam := pathParams["cineProjectId"]
	if cineProjectIdPathParam == ""{
		ErrorResponseWithMsg(http.StatusBadRequest, "request is missing required path parameter 'cineProjectId'", w)
		return
	}
	cineProjectId, err = ParseInt64(cineProjectIdPathParam)
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("DELETECineProjects_CineProjectId was unable to parse the path parameter 'cineProjectId', err: %s", err)
	}

	response, err := h.DELETECineProjects_CineProjectId(r.Context(), cineProjectId)
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("DELETECineProjects_CineProjectId returned err: %s", err)
	}

	if err = response.Send(w); err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("DELETECineProjects_CineProjectId was unable to send it's response, err: %s", err)
	}
}

// HandleGETCineProjects handles parsing input to pass to the GETCineProjects operation and sends responses back to the client
func (h *APIHandler) HandleGETCineProjects(w http.ResponseWriter, r *http.Request) {
	var err error
	response, err := h.GETCineProjects(r.Context())
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("GETCineProjects returned err: %s", err)
	}

	if err = response.Send(w); err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("GETCineProjects was unable to send it's response, err: %s", err)
	}
}

// HandleGETCineProjects_CineProjectId handles parsing input to pass to the GETCineProjects_CineProjectId operation and sends responses back to the client
func (h *APIHandler) HandleGETCineProjects_CineProjectId(w http.ResponseWriter, r *http.Request) {
	var err error
	pathParams := mux.Vars(r)

	var cineProjectId int64
	cineProjectIdPathParam := pathParams["cineProjectId"]
	if cineProjectIdPathParam == ""{
		ErrorResponseWithMsg(http.StatusBadRequest, "request is missing required path parameter 'cineProjectId'", w)
		return
	}
	cineProjectId, err = ParseInt64(cineProjectIdPathParam)
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("GETCineProjects_CineProjectId was unable to parse the path parameter 'cineProjectId', err: %s", err)
	}

	response, err := h.GETCineProjects_CineProjectId(r.Context(), cineProjectId)
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("GETCineProjects_CineProjectId returned err: %s", err)
	}

	if err = response.Send(w); err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("GETCineProjects_CineProjectId was unable to send it's response, err: %s", err)
	}
}

// HandlePOSTCineProjects handles parsing input to pass to the POSTCineProjects operation and sends responses back to the client
func (h *APIHandler) HandlePOSTCineProjects(w http.ResponseWriter, r *http.Request) {
	var err error
	reqBody := CineProject{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&reqBody); err != nil {
		ErrorResponseWithMsg(http.StatusBadRequest, "request body was not able to be parsed successfully 'CineProject'", w)
		return
	}
	if err := reqBody.Validate(); err != nil {
		errMsg := fmt.Errorf("request body was parsed successfully but failed validation, err: %w", err)
		ErrorResponseWithMsg(http.StatusBadRequest, errMsg.Error(), w)
		return
	}

	response, err := h.POSTCineProjects(r.Context(), reqBody)
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("POSTCineProjects returned err: %s", err)
	}

	if err = response.Send(w); err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("POSTCineProjects was unable to send it's response, err: %s", err)
	}
}

// HandlePUTCineProjects_CineProjectId handles parsing input to pass to the PUTCineProjects_CineProjectId operation and sends responses back to the client
func (h *APIHandler) HandlePUTCineProjects_CineProjectId(w http.ResponseWriter, r *http.Request) {
	var err error
	pathParams := mux.Vars(r)

	var cineProjectId int64
	cineProjectIdPathParam := pathParams["cineProjectId"]
	if cineProjectIdPathParam == ""{
		ErrorResponseWithMsg(http.StatusBadRequest, "request is missing required path parameter 'cineProjectId'", w)
		return
	}
	cineProjectId, err = ParseInt64(cineProjectIdPathParam)
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("PUTCineProjects_CineProjectId was unable to parse the path parameter 'cineProjectId', err: %s", err)
	}

	reqBody := CineProject{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&reqBody); err != nil {
		ErrorResponseWithMsg(http.StatusBadRequest, "request body was not able to be parsed successfully 'CineProject'", w)
		return
	}
	if err := reqBody.Validate(); err != nil {
		errMsg := fmt.Errorf("request body was parsed successfully but failed validation, err: %w", err)
		ErrorResponseWithMsg(http.StatusBadRequest, errMsg.Error(), w)
		return
	}

	response, err := h.PUTCineProjects_CineProjectId(r.Context(), cineProjectId, reqBody)
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("PUTCineProjects_CineProjectId returned err: %s", err)
	}

	if err = response.Send(w); err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("PUTCineProjects_CineProjectId was unable to send it's response, err: %s", err)
	}
}

