package model

// DataLayer is the interfacing layer between model logic and persistent data.
type DataLayer interface {
	AddUserRecord(u UserRecord)
	DeleteUserRecord(uid uint64)
	UserRecordByID(uid uint64) (UserRecord, bool)
	UIDExists(uid uint64) bool
}

var dataLayer DataLayer

// SetDataLayer sets the local data access layer for model logic.
func SetDataLayer(l DataLayer) {
	dataLayer = l
}
