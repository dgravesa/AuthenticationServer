package controller

import (
	"encoding/json"
	"net/http"

	"github.com/dgravesa/AuthenticationServer/model"
)

func loginHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		postLogin(w, r)
	case http.MethodDelete:
		fallthrough // TODO implement
	default:
		// TODO error
	}
}

func postLogin(w http.ResponseWriter, r *http.Request) {
	login, err := model.ParseUserLogin(&r.Form)

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
	enc := json.NewEncoder(w)
	enc.Encode(session)
}
