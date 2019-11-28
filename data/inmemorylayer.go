package data

import (
	"github.com/dgravesa/AuthenticationServer/model"
)

// InMemoryLayer provides a data store in memory.
type InMemoryLayer struct {
	users map[uint64]model.User
}

// NewInMemoryLayer returns a new InMemoryLayer.
func NewInMemoryLayer() *InMemoryLayer {
	layer := new(InMemoryLayer)
	layer.users = make(map[uint64]model.User)
	return layer
}

// AddUser adds a user to the data.
func (l *InMemoryLayer) AddUser(u model.User) {
	l.users[u.ID] = u
}

// DeleteUser removes a user from the data.
func (l *InMemoryLayer) DeleteUser(uid uint64) {
	delete(l.users, uid)
}

// UIDExists returns true if the questioned uid exists as a record in this data layer, otherwise false.
func (l *InMemoryLayer) UIDExists(uid uint64) bool {
	_, exists := l.users[uid]
	return exists
}
