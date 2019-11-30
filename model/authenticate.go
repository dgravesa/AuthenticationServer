package model

// AuthenticateUser returns a new session and true if the provided credentials are correct,
// otherwise a zeroed session and false.
func AuthenticateUser(login UserLogin) (UserSession, bool) {
	record, found := userRecordDataLayer.UserRecordByID(login.ID)
	if !found {
		return UserSession{}, false
	}

	loginHash := saltAndHash(login.Password, record.Salt)
	if loginHash != record.Hash {
		return UserSession{}, false
	}

	return newSession(login.ID), true
}
