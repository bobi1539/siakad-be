package service

import (
	"database/sql"
	"siakad/helper"
)

func BeginTransaction(db *sql.DB) *sql.Tx {
	tx, err := db.Begin()
	helper.PanicIfError(err)
	return tx
}
