package services

import (
	"database/sql"
	pb "user_service/genproto/user"
	"user_service/storage/postgres"
)

type UserService struct {
	pb.UnimplementedUserServer
	Repo *postgres.UserRepo
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{Repo: postgres.NewUserRepository(db)}
}

