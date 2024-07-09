package services

import (
	"context"
	"database/sql"
	pb "user_service/genproto/user"
	"user_service/storage/postgres"

	"github.com/pkg/errors"
)

type UserService struct {
	pb.UnimplementedUserServer
	Repo *postgres.UserRepo
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{Repo: postgres.NewUserRepository(db)}
}

func (s *UserService) GetUser(ctx context.Context, req *pb.ID) (*pb.UserInfo, error) {
	resp, err := s.Repo.GetUserByID(ctx, req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read user")
	}

	return resp, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UserInfo) (*pb.Void, error) {
	err := s.Repo.UpdateUser(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update user")
	}

	return &pb.Void{}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.ID) (*pb.Void, error) {
	err := s.Repo.DeleteUser(ctx, req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to delete user")
	}

	return &pb.Void{}, nil
}

func (s *UserService) ValidateUser(ctx context.Context, req *pb.ID) (*pb.Status, error) {
	resp, err := s.Repo.ValidateUser(ctx, req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to validate user")
	}

	return resp, nil
}
