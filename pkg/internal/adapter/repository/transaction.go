package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Transaction interface {
	Start(ctx context.Context) (*sqlx.Tx, error)
	Finish(ctx context.Context, tx *sqlx.Tx, err error) error
}

func (ri RepositoryImpl) Start(ctx context.Context) (*sqlx.Tx, error) {
	tx, err := ri.dbMaster.BeginTxx(ctx, nil)
	return tx, err
}

func (ri RepositoryImpl) Finish(ctx context.Context, tx *sqlx.Tx, errQuery error) error {
	if errQuery != nil {
		if errRollback := ri.rollback(tx); errRollback != nil {
			return errRollback
		}
		return errQuery
	}
	if err := ri.complete(tx); err != nil {
		if errRollback := ri.rollback(tx); errRollback != nil {
			return errRollback
		}
		return err
	}
	return nil
}

func (ri RepositoryImpl) complete(tx *sqlx.Tx) error {
	return tx.Commit()
}

func (ri RepositoryImpl) rollback(tx *sqlx.Tx) error {
	return tx.Rollback()
}
