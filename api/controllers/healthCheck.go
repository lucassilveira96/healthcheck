package controllers

import (
	"net/http"

	"github.com/lucassilveira96/healthcheck/api/responses"
	"github.com/lucassilveira96/healthcheck/api/services"
)

func (server *Server) GetAllDockers(w http.ResponseWriter, r *http.Request) {
	resp := services.GetAllDockers()

	responses.JSON(w, http.StatusOK, resp)
}
