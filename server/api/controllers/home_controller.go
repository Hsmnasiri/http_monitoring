package controllers

import (
	"net/http"
	"server/api/utils/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To The HTTP Monitoring  App")

}