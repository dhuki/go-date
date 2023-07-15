package repository

import (
	"context"

	"github.com/dhuki/go-date/pkg/internal/adapter/repository/model"
	"github.com/jmoiron/sqlx"
)

type RelationUser interface {
	UpsertRelationUser(ctx context.Context, tx *sqlx.Tx, relationUser model.RelationUser) (err error)
}

func (r RepositoryImpl) UpsertRelationUser(ctx context.Context, tx *sqlx.Tx, relationUser model.RelationUser) (err error) {
	query := `
		insert into relation_user("userId", "candidateId", "relationType", "createdAt", "updatedAt")
		VALUES(?, ?, ?, ?, ?)
		ON CONFLICT("userId", "candidateId") DO UPDATE 
			SET "relationType" = ?, "updatedAt" = ?;`

	_, err = tx.Exec(query, relationUser)
	return
}
