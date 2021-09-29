package kvstore

import (
	"database/sql"
)

type KVTxn struct {
	txn *sql.Tx
}
