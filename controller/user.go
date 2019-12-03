package controller

import (
	"net/http"

	"github.com/dgravesa/AuthenticationServer/model"
)

func userHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		postUser(w, r)
	case http.MethodDelete:
		deleteUser(w, r)
	case http.MethodPut:
		putUser(w, r)
	default:
		// TODO error
	}
}

func postUser(w http.ResponseWriter, r *http.Request) {
	user, err := model.ParseUserLogin(r.Form)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// TODO respond with error message
		return
	}

	if model.UIDExists(user.ID) {
		w.WriteHeader(http.StatusConflict)
		// TODO respond with error message
	} else {
		model.AddUserLogin(user)
		w.WriteHeader(http.StatusCreated)
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	uid, err := model.ParseUID(r.Form)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// TODO respond with error message
		return
	}

	if !model.UIDExists(uid) {
		w.WriteHeader(http.StatusNotFound)
		// TODO respond with error message
	} else {
		model.DeleteUserLogin(uid)
		w.WriteHeader(http.StatusOK)
	}
}

func putUser(w http.ResponseWriter, r *http.Request) {
	user, err := model.ParseUserLogin(r.Form)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// TODO respond with error message
		return
	}

	if !model.UIDExists(user.ID) {
		w.WriteHeader(http.StatusNotFound)
		// TODO respond with error message
	} else {
		model.UpdateUserLogin(user.ID, user.Password)
		w.WriteHeader(http.StatusOK)
	}
}
