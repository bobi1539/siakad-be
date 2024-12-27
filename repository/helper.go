package repository

import (
	"database/sql"
	"siakad/helper"
	"siakad/model/dto"
)

func FetchRows(repoCtx dto.RepositoryContext, sqlQuery string, args ...any) *sql.Rows {
	rows, err := repoCtx.Tx.QueryContext(repoCtx.Ctx, sqlQuery, args...)
	helper.PanicIfError(err)
	return rows
}
