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

// DeleteSession deletes a user session from the data store if it exists.
func (l *InMemorySessionLayer) DeleteSession(s model.Session) {
	for i, dataSession := range l.sessions {
		if s.UID == dataSession.UID && s.Key == dataSession.Key {
			l.sessions = append(l.sessions[:i], l.sessions[i+1:]...)
			return
		}
	}
}

// DeleteAllByUID removes all sessions associated with a particular user ID from the data store.
func (l *InMemorySessionLayer) DeleteAllByUID(uid uint64) {
	for i, session := range l.sessions {
		if uid == session.UID {
			l.sessions = append(l.sessions[:i], l.sessions[i+1:]...)
		}
	}
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
