package dto

import (
	"context"
	"database/sql"
)

type RepositoryContext struct {
	Ctx context.Context
	Tx  *sql.Tx
}

func BuildRepoCtx(ctx context.Context, tx *sql.Tx) RepositoryContext {
	return RepositoryContext{
		Ctx: ctx,
		Tx:  tx,
	}
}
