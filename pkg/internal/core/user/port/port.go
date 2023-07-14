package port

import (
	"context"
	"database/sql"

	modelReq "github.com/dhuki/go-date/pkg/internal/adapter/http/v1/model"
	modelRepo "github.com/dhuki/go-date/pkg/internal/adapter/repository/model"
)

type UserService interface {
	SignUp(ctx context.Context, request modelReq.CreateUserRequest) (err error)
	Login(ctx context.Context, request modelReq.LoginRequest) (response modelReq.LoginResponse, err error)
}

type UserRepository interface {
	Start(ctx context.Context) (*sql.Tx, error)
	Finish(ctx context.Context, tx *sql.Tx, errQuery error) error
	Create(ctx context.Context, tx *sql.Tx, user modelRepo.User) (id uint64, err error)
	GetUserByUsername(ctx context.Context, username string) (user modelRepo.User, err error)
	GetUserByEmail(ctx context.Context, email string) (user modelRepo.User, err error)
}
