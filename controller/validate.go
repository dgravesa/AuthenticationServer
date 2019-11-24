package controller

import "net/http"

func validateHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getValidate(w, r)
	default:
		// TODO error
	}
}

func getValidate(w http.ResponseWriter, r *http.Request) {
	// TODO implement
}
