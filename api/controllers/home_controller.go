package controllers

import (
	"http_monitoring/api/utils/responses"
	"net/http"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To The HTTP Monitoring  App")

}