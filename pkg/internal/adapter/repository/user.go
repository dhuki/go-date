package repository

import (
	"context"
	"database/sql"

	"github.com/dhuki/go-date/pkg/internal/adapter/repository/model"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user model.User) (id uint64, err error)
	GetUserByUsername(ctx context.Context, username string) (user model.User, err error)
	GetUserByEmail(ctx context.Context, email string) (user model.User, err error)
}

func (u RepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user model.User) (id uint64, err error) {
	query := `
		INSERT INTO user ("firstName", "lastName", email, "picUrl", district, city, "isPremium", "createdAt", "updatedAt") 
		VALUES (:firstName, :lastName, :email, :picUrl, :district, :city, :isPremium, :createdAt, :updatedAt)`

	result, err := tx.Exec(query, user)
	if err != nil {
		return
	}
	idInt, _ := result.LastInsertId()
	id = uint64(idInt)
	return
}

func (u RepositoryImpl) GetUserByUsername(ctx context.Context, username string) (user model.User, err error) {
	return
}

func (u RepositoryImpl) GetUserByEmail(ctx context.Context, email string) (user model.User, err error) {
	return
}
