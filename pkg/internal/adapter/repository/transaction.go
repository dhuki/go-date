package repository

import (
	"context"
	"database/sql"
)

type Transaction interface {
	Start(ctx context.Context) (*sql.Tx, error)
	Finish(ctx context.Context, tx *sql.Tx, errQuery error) error
}

func (ri RepositoryImpl) Start(ctx context.Context) (*sql.Tx, error) {
	tx, err := ri.dbMaster.BeginTx(ctx, nil)
	return tx, err
}

func (ri RepositoryImpl) Finish(ctx context.Context, tx *sql.Tx, errQuery error) error {
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

func (ri RepositoryImpl) complete(tx *sql.Tx) error {
	return tx.Commit()
}

func (ri RepositoryImpl) rollback(tx *sql.Tx) error {
	return tx.Rollback()
}
