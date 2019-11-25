package controller

import "net/http"

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
	// TODO implement
}
