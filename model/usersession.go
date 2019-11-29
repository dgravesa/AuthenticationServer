package model

import (
	"crypto/rand"
	"encoding/hex"
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
