package data

import "github.com/dgravesa/AuthenticationServer/model"

// InMemoryUserSessionLayer provides a data store in memory for user sessions.
type InMemoryUserSessionLayer struct {
	sessions []model.UserSession
}

// AddSession adds a user session to the data store.
func (l *InMemoryUserSessionLayer) AddSession(s model.UserSession) {
	l.sessions = append(l.sessions, s)
}
