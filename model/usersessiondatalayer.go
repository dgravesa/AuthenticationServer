package model

// UserSessionDataLayer is the interfacing layer between model logic and persistent data for user sessions.
type UserSessionDataLayer interface {
	AddSession(s UserSession)
	SessionExists(s UserSession) bool
}

var userSessionDataLayer UserSessionDataLayer

// SetUserSessionDataLayer sets the local data access layer for model logic.
func SetUserSessionDataLayer(l UserSessionDataLayer) {
	userSessionDataLayer = l
}
