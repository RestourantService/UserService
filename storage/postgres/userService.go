package postgres

import "github.com/jmoiron/sqlx"

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}
