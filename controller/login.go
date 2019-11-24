package controller

import "net/http"

func loginHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		postLogin(w, r)
	case http.MethodDelete:
		deleteLogin(w, r)
	default:
		// TODO error
	}
}

func postLogin(w http.ResponseWriter, r *http.Request) {
	// TODO implement
}

func deleteLogin(w http.ResponseWriter, r *http.Request) {
	// TODO implement
}
