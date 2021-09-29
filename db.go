package kvstore

import (
	"database/sql"
)

type Row interface {
	Scan(...interface{}) error
}

type DB interface {
	QueryRow(string, ...interface{}) *sql.Row
	Prepare(string) (*sql.Stmt, error)
}

func Read(db DB, key string) (string, error) {
	var value string
	err := db.QueryRow("select v from kvstore where k = ?", key).Scan(&value)
	return value, err
}

func Write(db DB, key string, value string) error {
	stmt, err := db.Prepare("replace into kvstore values (?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(key, value)
	if err != nil {
		return err
	}
	return nil
}
