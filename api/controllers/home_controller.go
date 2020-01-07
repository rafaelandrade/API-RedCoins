package controllers

import (
	"net/http"

	"github.com/rafaelandrade/API-RedCoins/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "RedCoins API")
}
