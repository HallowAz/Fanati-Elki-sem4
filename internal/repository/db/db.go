package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TxCommitter interface {
	InTx(ctx context.Context, fn func(ctx context.Context, tx pgx.Tx) error) error
}

type Committer struct {
	db *pgxpool.Pool
}

func NewTXCommitter(db *pgxpool.Pool) *Committer {
	return &Committer{
		db: db,
	}
}

func (c *Committer) InTx(ctx context.Context, fn func(ctx context.Context, tx pgx.Tx) error) (err error) {
	var tx pgx.Tx

	tx, err = c.db.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadWrite,
	})
	if err != nil {
		return err
	}

	// Обеспечиваем откат в случае ошибки
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback(ctx)

			panic(p)

			return
		}

		if err != nil {
			_ = tx.Rollback(ctx)

			return
		}

		err = tx.Commit(ctx)
	}()

	err = fn(ctx, tx)

	return
}
