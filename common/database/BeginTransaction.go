package database

import (
	"context"
	"database/sql"
)

func BeginTransaction(context context.Context, db *sql.DB) (response *sql.Tx, error error) {
	tx, err := db.BeginTx(context, nil)
	return tx, err
}
