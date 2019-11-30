package model

// AuthenticateUser returns a new session and true if the provided credentials are correct,
// otherwise a zeroed session and false.
func AuthenticateUser(login UserLogin) (Session, bool) {
	record, found := userRecordDataLayer.UserRecordByID(login.ID)
	if !found {
		return Session{}, false
	}

	loginHash := applySaltAndHash(login.Password, record.Salt)
	if loginHash != record.Hash {
		return Session{}, false
	}

	session := newSession(login.ID)
	sessionDataLayer.AddSession(session)
	return session, true
}
