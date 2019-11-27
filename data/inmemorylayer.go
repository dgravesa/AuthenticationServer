package data

import "github.com/dgravesa/AuthenticationServer/model"

// InMemoryLayer provides a data store in memory.
type InMemoryLayer struct {
	users []model.User
}

// NewInMemoryLayer returns a new InMemoryLayer.
func NewInMemoryLayer() *InMemoryLayer {
	return new(InMemoryLayer)
}

// AddUser adds a user to the data.
func (l *InMemoryLayer) AddUser(u model.User) {
	l.users = append(l.users, u)
}

// UIDExists returns true if the questioned uid exists as a record in this data layer, otherwise false.
func (l *InMemoryLayer) UIDExists(uid uint64) bool {
	for _, u := range l.users {
		if u.ID == uid {
			return true
		}
	}
	return false
}
