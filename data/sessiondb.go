package data

import (
	"database/sql"

	"github.com/dgravesa/AuthenticationServer/model"
	"github.com/dgravesa/dbcommon/dbserver"
)

const sessionDatabaseName = "sessions"

// SessionDB is an interface to a session database.
type SessionDB struct {
	db *sql.DB
}

// NewSessionDB creates a new session database layer with the database server configuration specified by cfg.
func NewSessionDB(cfg dbserver.Config) (*SessionDB, error) {
	var sessionDB *SessionDB

	db, err := dbserver.StartupDB(cfg, databaseName)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(sessionDBCreateTableIfDoesNotExistQuery)
	if err != nil {
		sessionDB = &SessionDB{db: db}
	}

	return sessionDB, err
}

// AddSession adds s to the database.
func AddSession(s model.Session) {
	// TODO implement
}

// SessionExists returns true if s exists in the database, otherwise false.
func SessionExists(s model.Session) bool {
	// TODO implement
	return false
}

// DeleteSession deletes s if it exists in the database.
func DeleteSession(s model.Session) {
	// TODO implement
}

// DeleteAllByUID deletes all sessions associated with uid.
func DeleteAllByUID(uid uint64) {
	// TODO implement
}
