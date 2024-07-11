package service

import (
	"context"
	"errors"
	"log"
	"user_service/api/auth"
	pb "user_service/genproto/authentication"
)

func (S *UserService) Register(ctx context.Context, req *pb.UserDetails) (*pb.ID, error) {
	return S.Repo.Register(ctx, req)
}

func (S *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	checker, err := S.Repo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	if req.Password != checker.Password {
		return nil, errors.New("username or password error")
	}

	res := &pb.LoginResponse{
		Access: &pb.AccessToken{
			Id:       checker.Id,
			Username: checker.Username,
			Email:    checker.Email,
		},
		Refresh: &pb.RefreshToken{
			Userid: checker.Id,
		},
	}
	err = auth.GeneratedRefreshJWTToken(res)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	err = auth.GeneratedAccessJWTToken(res)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	err = S.Repo.StoreRefreshToken(ctx, res)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return res, nil
}

func (S *UserService) Refresh(ctx context.Context, req *pb.CheckRefreshTokenRequest) (*pb.CheckRefreshTokenResponse, error) {
	_, err := auth.ValidateRefreshToken(req.Token)
	if err != nil {
		log.Print(err)
		return &pb.CheckRefreshTokenResponse{Acces: false}, err
	}
	id, err := auth.GetUserIdFromRefreshToken(req.Token)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	info, err := S.Repo.GetUserByID(ctx, id)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	res := pb.LoginResponse{
		Access: &pb.AccessToken{
			Id:       info.Id,
			Username: info.Username,
			Email:    info.Email,
		},
	}
	err = auth.GeneratedAccessJWTToken(&res)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return &pb.CheckRefreshTokenResponse{Acces: true, Accestoken: res.Access.Accesstoken}, nil
}

