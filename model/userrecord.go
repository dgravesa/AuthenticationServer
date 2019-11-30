package model

// UserRecord contains information stored in the data for authenticating a user.
type UserRecord struct {
	ID   uint64
	Salt string
	Hash string
}

func makeUserRecord(ul UserLogin) UserRecord {
	salt := makeSalt()
	hash := applySaltAndHash(ul.Password, salt)
	return UserRecord{ul.ID, salt, hash}
}
