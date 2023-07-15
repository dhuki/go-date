package port

import (
	"context"
	"time"

	modelReq "github.com/dhuki/go-date/pkg/internal/adapter/http/v1/model"
	modelRepo "github.com/dhuki/go-date/pkg/internal/adapter/repository/model"
	"github.com/jmoiron/sqlx"
)

type CandidateService interface {
	SwipeAction(ctx context.Context, candidateID uint64, swipeDirection string) (err error)
	GetListCandidate(ctx context.Context, limit int) (resp modelReq.CandidateListPaginationReponse, err error)
}

type CandidateRepository interface {
	Start(ctx context.Context) (*sqlx.Tx, error)
	Finish(ctx context.Context, tx *sqlx.Tx, err error) error
	UpsertRelationUser(ctx context.Context, tx *sqlx.Tx, relationUser modelRepo.RelationUser) (err error)
	GetUserByID(ctx context.Context, id uint64) (user modelRepo.User, err error)
	GetUserPagination(ctx context.Context, gender string, limit, offset int) (users []modelRepo.User, err error)
}

type RedisLibs interface {
	Set(key string, value interface{}, ttl time.Duration) (err error)
	Get(key string) (value string)
}
