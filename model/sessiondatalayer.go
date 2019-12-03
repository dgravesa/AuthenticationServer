package model

// SessionDataLayer is the interfacing layer between model logic and persistent data for user sessions.
type SessionDataLayer interface {
	AddSession(s Session)
	SessionExists(s Session) bool
	DeleteSession(s Session)
	DeleteAllByUID(uid uint64)
}

var sessionDataLayer SessionDataLayer

// SetSessionDataLayer sets the local data access layer for model logic.
func SetSessionDataLayer(l SessionDataLayer) {
	sessionDataLayer = l
}
