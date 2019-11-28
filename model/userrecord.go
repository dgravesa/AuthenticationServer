package model

// UserRecord contains information stored in the data for authenticating a user.
type UserRecord struct {
	uid  uint64
	salt string
	hash string
}

func makeUserRecord(ul UserLogin) UserRecord {
	salt := makeSalt()
	hash := saltAndHash(ul.Password, salt)
	return UserRecord{ul.ID, salt, hash}
}
