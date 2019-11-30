package model

// UserRecordDataLayer is the interfacing layer between model logic and persistent data for user records.
type UserRecordDataLayer interface {
	AddUserRecord(u UserRecord)
	DeleteUserRecord(uid uint64)
	UserRecordByID(uid uint64) (UserRecord, bool)
	UIDExists(uid uint64) bool
}

var userRecordDataLayer UserRecordDataLayer

// SetUserRecordDataLayer sets the local data access layer for model logic.
func SetUserRecordDataLayer(l UserRecordDataLayer) {
	userRecordDataLayer = l
}
