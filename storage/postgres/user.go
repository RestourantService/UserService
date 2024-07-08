package postgres

import (
	"database/sql"
	pb "user_service/genproto/user"
)


type UserRepo struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (u *UserRepo) GetUserByID(id string) (pb.UserInfo, error) {
	var user pb.UserInfo
	u.DB.Exec("select ")
}