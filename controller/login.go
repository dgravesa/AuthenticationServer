package controller

import (
	"net/http"

	"github.com/dgravesa/AuthenticationServer/model"
)

func loginHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		postLogin(w, r)
	case http.MethodDelete:
		deleteLogin(w, r) // logout
	default:
		// TODO error
	}
}

func postLogin(w http.ResponseWriter, r *http.Request) {
	login, err := model.ParseUserLogin(r.Form)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// TODO respond with error message
		return
	}

	session, loginSucceeded := model.AuthenticateUser(login)

	if !loginSucceeded {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	model.EncodeSessionToJSON(w, session)
}

func deleteLogin(w http.ResponseWriter, r *http.Request) {
	session, err := model.ParseSession(r.URL.Query())

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !model.SessionExists(session) {
		w.WriteHeader(http.StatusNotFound)
	} else {
		model.DeleteSession(session)
		w.WriteHeader(http.StatusOK)
	}
}
