package data

const sessionDBCreateTableIfDoesNotExistQuery string = `
CREATE TABLE IF NOT EXISTS sessions (
	id SERIAL PRIMARY KEY NOT NULL,
	uid BIGINT NOT NULL,
	key CHAR(64) NOT NULL
);`
