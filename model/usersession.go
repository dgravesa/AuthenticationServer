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
	// TODO implement
	return UserSession{}, fmt.Errorf("not yet implemented")
}
