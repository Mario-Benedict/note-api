package controllers

import (
	"net/http"

	"github.com/Mario-Benedict/note-api/helper"
)

type HealthController struct{}

func (HealthController) Check(res http.ResponseWriter, req *http.Request) {
	helper.WriteJSON(res, map[string]string{
		"status": "ok",
	}, http.StatusOK)
}
