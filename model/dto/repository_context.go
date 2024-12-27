package dto

import (
	"context"
	"database/sql"
)

type RepositoryContext struct {
	Ctx context.Context
	Tx  *sql.Tx
}
