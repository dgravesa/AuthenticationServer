package model

// DataLayer is the interfacing layer between model logic and persistent data.
type DataLayer interface {
	AddUser(u UserLogin)
	DeleteUser(uid uint64)
	UIDExists(uid uint64) bool
}

var dataLayer DataLayer

// SetDataLayer sets the local data access layer for model logic.
func SetDataLayer(l DataLayer) {
	dataLayer = l
}
