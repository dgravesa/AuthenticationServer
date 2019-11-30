package model

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
)

const sessionKeyLen = 32

// UserSession is used for subsequent authentication when a user has successfully logged in.
type UserSession struct {
	UID uint64 `json:"userId"`
	Key string `json:"key"`
}

// SessionExists returns true if the session is found in the data, otherwise false.
func SessionExists(s UserSession) bool {
	return userSessionDataLayer.SessionExists(s)
}

func makeSessionKey() string {
	keyBytes := make([]byte, sessionKeyLen)
	rand.Read(keyBytes)
	return hex.EncodeToString(keyBytes)
}

func newSession(uid uint64) UserSession {
	return UserSession{
		UID: uid,
		Key: makeSessionKey(),
	}
}

// EncodeSessionToJSON writes a JSON-encoded session to w.
func EncodeSessionToJSON(w io.Writer, s UserSession) {
	enc := json.NewEncoder(w)
	enc.Encode(s)
}

// DecodeSessionFromJSON reads a JSON-encoded session from r.
func DecodeSessionFromJSON(r io.Reader) (UserSession, error) {
	var session UserSession

	var nillable struct {
		UID *uint64 `json:"userId"`
		Key *string `json:"key"`
	}

	dec := json.NewDecoder(r)
	if err := dec.Decode(&nillable); err != nil {
		return session, err
	} else if nillable.UID == nil {
		return session, fmt.Errorf("session JSON: missing \"userId\"")
	} else if nillable.Key == nil {
		return session, fmt.Errorf("session JSON: missing \"key\"")
	}

	session.UID = *nillable.UID
	session.Key = *nillable.Key
	return session, nil
}
