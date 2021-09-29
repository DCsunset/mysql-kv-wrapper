package kvstore

import (
	"database/sql"
)

type KVTxn struct {
	txn *sql.Tx
}

func (t *KVTxn) Read(key string) (string, error) {
	return Read(t.txn, key)
}

func (t *KVTxn) Write(key string, value string) error {
	return Write(t.txn, key, value)
}

func (t *KVTxn) Commit() error {
	return t.txn.Commit()
}

func (t *KVTxn) Rollback() error {
	return t.txn.Rollback()
}
