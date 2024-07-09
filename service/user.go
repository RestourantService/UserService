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

func (s *UserService) GetUser(ctx context.Context, req *pb.ID) (*pb.UserInfo, error) {
	return s.Repo.GetUserByID(ctx, req.Id)
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UserInfo) (*pb.Void, error) {
	return &pb.Void{}, s.Repo.UpdateUser(ctx, req)
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.ID) (*pb.Void, error) {
	return &pb.Void{}, s.Repo.DeleteUser(ctx, req.Id)
}

func (s *UserService) ValidateUser(ctx context.Context, req *pb.ID) (*pb.Status, error) {
	return s.Repo.ValidateUser(ctx, req.Id)
}
