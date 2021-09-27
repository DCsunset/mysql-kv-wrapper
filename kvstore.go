package kvstore

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type KVStore struct {
	db *sql.DB
}

func (s *KVStore) Open(source string) error {
	db, err := sql.Open("mysql", source)
	if err != nil {
		s.db = nil
		return err
	}
	s.db = db
	_, err = db.Exec("create table if not exists kvstore (key varchar(32), value varchar(1024), primary key (key))")
	if err != nil {
		return err
	}
	return nil
}

func (s *KVStore) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}

func (s *KVStore) Read(key string) (string, error) {
	var value string
	err := s.db.QueryRow("select value from kvstore where key = ?", key).Scan(&value)
	return value, err
}

func (s *KVStore) Write(key string, value string) error {
	stmt, err := s.db.Prepare("replace into kvstore values (?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(value)
	if err != nil {
		return err
	}
	return nil
}
