package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/dhuki/go-date/pkg/internal/adapter/repository/model"
	"github.com/jmoiron/sqlx"
)

type RelationUser interface {
	GetRelationUserByUserIdAndCandidate(ctx context.Context, userID, candidateID uint64) (relationUser model.RelationUser, err error)
	UpsertRelationUser(ctx context.Context, tx *sqlx.Tx, relationUser model.RelationUser) (err error)
}

func (r RepositoryImpl) GetRelationUserByUserIdAndCandidate(ctx context.Context, userID, candidateID uint64) (relationUser model.RelationUser, err error) {
	return relationUser, r.dbSlave.GetContext(ctx, &relationUser, `select * from relation_users where "userId" = $1 and "candidateId" = $2 and "deletedAt" is null order by id desc limit 1`, userID, candidateID)
}

func (r RepositoryImpl) UpsertRelationUser(ctx context.Context, tx *sqlx.Tx, relationUser model.RelationUser) (err error) {
	relationUser.UpdatedAt = time.Now()

	arrDefaultColumn := []string{`"userId"`, `"candidateId"`, `"relationType"`, `"createdAt"`, `"updatedAt"`}
	if relationUser.ID > 0 {
		arrDefaultColumn = append([]string{"id"}, arrDefaultColumn...)
	} else {
		relationUser.CreatedAt = time.Now()
	}

	var arrPlaceholder []string
	for _, v := range arrDefaultColumn {
		arrPlaceholder = append(arrPlaceholder, fmt.Sprintf(":%s", strings.ReplaceAll(v, `"`, ``)))
	}

	placeholder, defaultColumn := strings.Join(arrPlaceholder, ","), strings.Join(arrDefaultColumn, ",")
	query := `
		insert into relation_users(%s)
		values(%s)
		on conflict("id") do update 
			set "relationType" = :relationType, "updatedAt" = :updatedAt`

	_, err = tx.NamedExecContext(ctx, fmt.Sprintf(query, defaultColumn, placeholder), relationUser)
	return
}
