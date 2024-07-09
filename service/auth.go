package service

import (
	"context"
	"errors"
	pb "user_service/genproto/authentication"
)

func (S *UserService) Register(ctx context.Context, req *pb.UserDetails) (*pb.ID, error) {
	return S.Repo.Register(ctx, req)
}

func (S *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.ID, error) {
	checker, err := S.Repo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if req.Password != checker.Password {
		return nil, errors.New("username or password error")
	}
	return &pb.ID{Id: checker.Id}, nil
}
func (S *UserService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.Void, error) {
	checker, err := S.Repo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	err = S.Repo.DeleteRefreshToken(ctx, checker.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (S *UserService) RefreshToken(ctx context.Context, req *pb.TokenRequest) (*pb.Void, error) {
	return &pb.Void{}, S.Repo.StoreRefreshToken(ctx, req)
}
