package data

import (
	"github.com/dgravesa/AuthenticationServer/model"
)

// InMemoryLayer provides a data store in memory.
type InMemoryLayer struct {
	users map[uint64]model.UserRecord
}

// NewInMemoryLayer returns a new InMemoryLayer.
func NewInMemoryLayer() *InMemoryLayer {
	layer := new(InMemoryLayer)
	layer.users = make(map[uint64]model.UserRecord)
	return layer
}

// AddUserRecord adds a user record to the data.
func (l *InMemoryLayer) AddUserRecord(u model.UserRecord) {
	l.users[u.ID] = u
}

// DeleteUserRecord removes a user record from the data.
func (l *InMemoryLayer) DeleteUserRecord(uid uint64) {
	delete(l.users, uid)
}

// UserRecordByID returns the corresponding user record and true if the ID exists,
// otherwise a zeroed UserRecord and false.
func (l *InMemoryLayer) UserRecordByID(uid uint64) (model.UserRecord, bool) {
	record, exists := l.users[uid]
	return record, exists
}

// UIDExists returns true if the questioned uid exists as a record in this data layer, otherwise false.
func (l *InMemoryLayer) UIDExists(uid uint64) bool {
	_, exists := l.users[uid]
	return exists
}
