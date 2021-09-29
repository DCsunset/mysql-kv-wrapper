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

	_, err = db.Exec("create database if not exists kvstore")
	if err != nil {
		return err
	}

	_, err = db.Exec("use kvstore")
	if err != nil {
		return err
	}

	_, err = db.Exec("create table if not exists kvstore (k varchar(32), v varchar(1024), primary key (k))")
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

func (s *KVStore) Begin() (*KVTxn, error) {
	txn, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	return &KVTxn{txn}, nil
}

func (s *KVStore) Read(key string) (string, error) {
	return Read(s.db, key)
}

func (s *KVStore) Write(key string, value string) error {
	return Write(s.db, key, value)
}
