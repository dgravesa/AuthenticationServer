package data

import (
	"github.com/dgravesa/AuthenticationServer/model"
)

// InMemoryUserRecordLayer provides a data store in memory for user records.
type InMemoryUserRecordLayer struct {
	users map[uint64]model.UserRecord
}

// NewInMemoryUserRecordLayer returns a new InMemoryUserRecordLayer.
func NewInMemoryUserRecordLayer() *InMemoryUserRecordLayer {
	layer := new(InMemoryUserRecordLayer)
	layer.users = make(map[uint64]model.UserRecord)
	return layer
}

// AddUserRecord adds a user record to the data.
func (l *InMemoryUserRecordLayer) AddUserRecord(u model.UserRecord) {
	l.users[u.ID] = u
}

// DeleteUserRecord removes a user record from the data.
func (l *InMemoryUserRecordLayer) DeleteUserRecord(uid uint64) {
	delete(l.users, uid)
}

// UpdateUserRecord updates a user record if it exists.
func (l *InMemoryUserRecordLayer) UpdateUserRecord(u model.UserRecord) {
	if l.UIDExists(u.ID) {
		l.users[u.ID] = u
	}
}

// UserRecordByID returns the corresponding user record and true if the ID exists,
// otherwise a zeroed UserRecord and false.
func (l *InMemoryUserRecordLayer) UserRecordByID(uid uint64) (model.UserRecord, bool) {
	record, exists := l.users[uid]
	return record, exists
}

// UIDExists returns true if the questioned uid exists as a record in this data layer, otherwise false.
func (l *InMemoryUserRecordLayer) UIDExists(uid uint64) bool {
	_, exists := l.users[uid]
	return exists
}
