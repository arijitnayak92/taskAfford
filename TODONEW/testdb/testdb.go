package testdb

import (
	"database/sql"
	"os"
)

const createTable = `
DROP TABLE IF EXISTS todo;
Alter SEQUENCE todo_id RESTART WITH 1;
CREATE TABLE IF NOT EXISTS todo (
  ID int default nextval('todo_id'::regclass),
  TITLE TEXT NOT NULL,
  NOTE TEXT,
  STATUS BOOLEAN
);
`

// const createTable = `
// CREATE TABLE IF NOT EXISTS todo (
//   ID int default nextval('todo_id'::regclass),
//   TITLE TEXT NOT NULL,
//   NOTE TEXT,
//   STATUS BOOLEAN
// );
// `

type TestDB struct {
	db *sql.DB
}

func Setup() *sql.DB {
	db, err := connectPostgresForTests()
	if err != nil {
		panic(err)
	}

	if _, err = db.Exec(createTable); err != nil {
		panic(err)
	}

	return db
}

func connectPostgresForTests() (*sql.DB, error) {
	connStr := os.Getenv("POSTGRES_URI")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
