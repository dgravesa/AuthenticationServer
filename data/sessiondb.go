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
func (l *SessionDB) AddSession(s model.Session) {
	l.db.Exec(sessionDBInsertSessionQuery, s.UID, s.Key)
}

// SessionExists returns true if s exists in the database, otherwise false.
func (l *SessionDB) SessionExists(s model.Session) bool {
	var found bool
	l.db.QueryRow(sessionDBFindSessionQuery, s.UID, s.Key).Scan(&found)
	return found
}

// DeleteSession deletes s if it exists in the database.
func (l *SessionDB) DeleteSession(s model.Session) {
	l.db.Exec(sessionDBDeleteSessionQuery, s.UID, s.Key)
}

// DeleteAllByUID deletes all sessions associated with uid.
func (l *SessionDB) DeleteAllByUID(uid uint64) {
	l.db.Exec(sessionDBDeleteAllSessionsByUIDQuery, uid)
}
