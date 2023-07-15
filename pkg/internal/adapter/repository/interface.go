package repository

import (
	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -destination=mocks/mock_repo.go -package=mocks github.com/dhuki/go-date/pkg/internal/adapter/repository Repository

type Repository interface {
	Transaction
	UserRepository
	RelationUser
}

type RepositoryImpl struct {
	dbMaster, dbSlave *sqlx.DB
}

func NewRepository(dbMaster, dbSlave *sqlx.DB) Repository {
	return RepositoryImpl{
		dbMaster: dbMaster,
		dbSlave:  dbSlave,
	}
}
