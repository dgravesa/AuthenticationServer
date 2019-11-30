package model

// UserRecordDataLayer is the interfacing layer between model logic and persistent data.
type UserRecordDataLayer interface {
	AddUserRecord(u UserRecord)
	DeleteUserRecord(uid uint64)
	UserRecordByID(uid uint64) (UserRecord, bool)
	UIDExists(uid uint64) bool
}

var dataLayer UserRecordDataLayer

// SetUserRecordDataLayer sets the local data access layer for model logic.
func SetUserRecordDataLayer(l UserRecordDataLayer) {
	dataLayer = l
}
