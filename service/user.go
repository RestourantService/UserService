package services

import (
	"context"
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

func (S *UserService) GetUser(ctx context.Context, req *pb.ID) (*pb.UserInfo, error) {
	return S.Repo.GetUserByID(ctx, req.Id)
}

func (S *UserService) UpdateUser(ctx context.Context, req *pb.UserInfo) (*pb.Void, error) {
	return &pb.Void{}, S.Repo.UpdateUser(ctx, req)
}

func (S *UserService) DeleteUser(ctx context.Context, req *pb.ID) (*pb.Void, error) {
	return &pb.Void{}, S.Repo.DeleteUser(ctx, req.Id)
}

func (S *UserService) ValidateUser(ctx context.Context, req *pb.ID) (*pb.Status, error) {
	return S.Repo.ValidateUser(ctx, req.Id)
}
