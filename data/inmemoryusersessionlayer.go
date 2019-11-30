package data

import "github.com/dgravesa/AuthenticationServer/model"

// InMemoryUserSessionLayer provides a data store in memory for user sessions.
type InMemoryUserSessionLayer struct {
	sessions []model.UserSession
}

// NewInMemoryUserSessionLayer returns a new in-memory user sessions data layer.
func NewInMemoryUserSessionLayer() *InMemoryUserSessionLayer {
	return new(InMemoryUserSessionLayer)
}

// AddSession adds a user session to the data store.
func (l *InMemoryUserSessionLayer) AddSession(s model.UserSession) {
	l.sessions = append(l.sessions, s)
}
