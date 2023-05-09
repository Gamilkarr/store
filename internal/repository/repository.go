package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type PgxIface interface {
	Begin(context.Context) (pgx.Tx, error)
	Close()
}

type Repository struct {
	Conn PgxIface
}

const isStoreAvailable = "SELECT availability FROM stores WHERE id=$1"
