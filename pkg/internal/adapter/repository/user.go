package repository

import (
	"context"
	"time"

	"github.com/dhuki/go-date/pkg/internal/adapter/repository/model"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sqlx.Tx, user model.User) (id uint64, err error)
	GetUserByID(ctx context.Context, id uint64) (user model.User, err error)
	GetUserByUsername(ctx context.Context, username string) (user model.User, err error)
	GetUserPagination(ctx context.Context, userID uint64, gender string, limit, offset int) (users []model.User, err error)
	CountTotalUserPagination(ctx context.Context, userID uint64, gender string) (count int, err error)
}

func (u RepositoryImpl) Create(ctx context.Context, tx *sqlx.Tx, user model.User) (id uint64, err error) {
	query := `
		insert into users ("username", "password", "firstName", "lastName", "gender", "picUrl", district, city, "isPremium", "createdAt", "updatedAt") 
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	result, err := tx.ExecContext(ctx, query, user.Username, user.Password, user.FirstName, user.LastName, user.Gender, user.PicUrl, user.District, user.City, false, time.Now(), time.Now())
	if err != nil {
		return
	}
	idInt, _ := result.LastInsertId()
	id = uint64(idInt)
	return
}

func (u RepositoryImpl) GetUserByID(ctx context.Context, id uint64) (user model.User, err error) {
	return user, u.dbSlave.QueryRowxContext(ctx, `
		select * from users where id = $1 and "deletedAt" is null;`,
		id).StructScan(&user)
}

func (u RepositoryImpl) GetUserByUsername(ctx context.Context, username string) (user model.User, err error) {
	return user, u.dbSlave.GetContext(ctx, &user, `select * from users where username = $1 and "deletedAt" is null`, username)
}

func (u RepositoryImpl) GetUserPagination(ctx context.Context, userID uint64, gender string, limit, offset int) (users []model.User, err error) {
	return users, u.dbSlave.SelectContext(ctx, &users, `
		select * from users 
		where "deletedAt" is null and id != $1 and gender != $2
		limit $3 offset $4`, userID, gender, limit, offset)
}

func (u RepositoryImpl) CountTotalUserPagination(ctx context.Context, userID uint64, gender string) (count int, err error) {
	return count, u.dbSlave.GetContext(ctx, &count, `
		select count(1) from users 
		where "deletedAt" is null and id != $1 and gender != $2`, userID, gender)
}
