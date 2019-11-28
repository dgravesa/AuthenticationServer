package model

// DataLayer is the interfacing layer between model logic and persistent data.
type DataLayer interface {
	AddUser(u User)
	DeleteUser(uid uint64)
	UIDExists(uid uint64) bool
}

var dataLayer DataLayer

// SetDataLayer sets the local data access layer for model logic.
func SetDataLayer(l DataLayer) {
	dataLayer = l
}

// AddUser adds a user to the data.
func AddUser(u User) {
	dataLayer.AddUser(u)
}

// DeleteUser removes the user associated with the ID from the data.
func DeleteUser(uid uint64) {
	dataLayer.DeleteUser(uid)
}

// UIDExists returns true if the questioned UID already exists in the data.
func UIDExists(uid uint64) bool {
	return dataLayer.UIDExists(uid)
}
