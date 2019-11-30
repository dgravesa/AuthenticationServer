package controller

import (
	"net/http"

	"github.com/dgravesa/AuthenticationServer/model"
)

func validateHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getValidate(w, r)
	default:
		// TODO error
	}
}

func getValidate(w http.ResponseWriter, r *http.Request) {
	session, err := model.ParseUserSession(r.URL.Query())

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if model.SessionExists(session) {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
