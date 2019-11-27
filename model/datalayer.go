package model

// DataLayer is the interfacing layer between model logic and persistent data.
type DataLayer interface {
	AddUser(u User)
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

// UIDExists returns true if the questioned UID already exists in the data.
func UIDExists(uid uint64) bool {
	return dataLayer.UIDExists(uid)
}
