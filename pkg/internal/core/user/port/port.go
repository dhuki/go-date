package port

import (
	"context"
	"time"

	modelReq "github.com/dhuki/go-date/pkg/internal/adapter/http/v1/model"
	modelRepo "github.com/dhuki/go-date/pkg/internal/adapter/repository/model"
	"github.com/jmoiron/sqlx"
)

type UserService interface {
	SignUp(ctx context.Context, request modelReq.CreateUserRequest) (err error)
	Login(ctx context.Context, request modelReq.LoginRequest) (response modelReq.LoginResponse, err error)
}

type UserRepository interface {
	Start(ctx context.Context) (*sqlx.Tx, error)
	Finish(ctx context.Context, tx *sqlx.Tx, err error) error
	Create(ctx context.Context, tx *sqlx.Tx, user modelRepo.User) (id uint64, err error)
	GetUserByUsername(ctx context.Context, username string) (user modelRepo.User, err error)
}

type JWTAccessToken interface {
	GenerateJWTAccessToken(userID uint64) (token string, err error)
}

type RedisLibs interface {
	Set(key string, value interface{}, ttl time.Duration) (err error)
	Get(key string) (value string)
	Delete(key string) (err error)
	SetIncr(key string) (count int64)
}
