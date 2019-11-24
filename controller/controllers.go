package controller

import "net/http"

// RegisterRoutes registers the handlers on their designated routes.
func RegisterRoutes() {
	http.HandleFunc("/user", userHandleFunc)
	http.HandleFunc("/login", loginHandleFunc)
	http.HandleFunc("/validate", validateHandleFunc)
}
