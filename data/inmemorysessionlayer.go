package data

import "github.com/dgravesa/AuthenticationServer/model"

// InMemorySessionLayer provides a data store in memory for user sessions.
type InMemorySessionLayer struct {
	sessions []model.Session
}

// NewInMemorySessionLayer returns a new in-memory user sessions data layer.
func NewInMemorySessionLayer() *InMemorySessionLayer {
	return new(InMemorySessionLayer)
}

// AddSession adds a user session to the data store.
func (l *InMemorySessionLayer) AddSession(s model.Session) {
	l.sessions = append(l.sessions, s)
}

// SessionExists returns true if the session is found in the data, otherwise false.
func (l *InMemorySessionLayer) SessionExists(s model.Session) bool {
	for _, dataSession := range l.sessions {
		if s.UID == dataSession.UID && s.Key == dataSession.Key {
			return true
		}
	}
	return false
}
