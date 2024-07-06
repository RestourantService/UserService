package services

import (
	userPb "userService/genproto/UserService"
	"userService/storage/postgres"

	"github.com/jmoiron/sqlx"
)

type userService struct {
	UserRepo *postgres.UserRepo
	userPb.UnimplementedUserServiceServer
}

func NewUserService(db *sqlx.DB) *userService {
	return &userService{UserRepo: postgres.NewUserRepository(db)}
}
