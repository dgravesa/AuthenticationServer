package data

const sessionDBCreateTableIfDoesNotExistQuery string = `
CREATE TABLE IF NOT EXISTS sessions (
	id SERIAL PRIMARY KEY NOT NULL,
	uid BIGINT NOT NULL,
	key CHAR(64) NOT NULL
);`

const sessionDBInsertSessionQuery string = `
INSERT INTO sessions (uid, key)
VALUES ($1, $2)`

const sessionDBDeleteSessionQuery string = `
DELETE FROM sessions
WHERE uid = $1 AND key = $2`

const sessionDBDeleteAllSessionsByUIDQuery string = `
DELETE FROM sessions
WHERE uid = $1`

const sessionDBFindSessionQuery string = `
SELECT EXISTS(
	SELECT uid, key
	FROM sessions
	WHERE uid = $1 AND key = $2
);`
