package model

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strconv"
)

const sessionKeyLen = 32

// Session is used for subsequent authentication when a user has successfully logged in.
type Session struct {
	UID uint64 `json:"userId"`
	Key string `json:"key"`
}

// SessionExists returns true if the session is found in the data, otherwise false.
func SessionExists(s Session) bool {
	return sessionDataLayer.SessionExists(s)
}

// DeleteSession removes a session from the data.
func DeleteSession(s Session) {
	sessionDataLayer.DeleteSession(s)
}

// ParseSession extracts a Session from http request query parameters.
func ParseSession(v url.Values) (Session, error) {
	var s Session
	var err error

	if s.UID, err = strconv.ParseUint(v.Get("userId"), 10, 64); err != nil {
		return Session{}, err
	}

	if s.Key = v.Get("key"); s.Key == "" {
		return Session{}, fmt.Errorf("No key given")
	}

	return s, nil
}

func makeSessionKey() string {
	keyBytes := make([]byte, sessionKeyLen)
	rand.Read(keyBytes)
	return hex.EncodeToString(keyBytes)
}

func newSession(uid uint64) Session {
	return Session{
		UID: uid,
		Key: makeSessionKey(),
	}
}

// EncodeSessionToJSON writes a JSON-encoded session to w.
func EncodeSessionToJSON(w io.Writer, s Session) {
	enc := json.NewEncoder(w)
	enc.Encode(s)
}

// DecodeSessionFromJSON reads a JSON-encoded session from r.
func DecodeSessionFromJSON(r io.Reader) (Session, error) {
	var session Session

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
