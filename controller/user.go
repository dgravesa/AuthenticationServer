package controller

import "net/http"

func userHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		postUser(w, r)
	case http.MethodDelete:
		deleteUser(w, r)
	case http.MethodPut:
		fallthrough // TODO implement
	default:
		// TODO error
	}
}

func postUser(w http.ResponseWriter, r *http.Request) {
	// TODO implement
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	// TODO implement
}
