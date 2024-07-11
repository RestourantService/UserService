package service

import (
	"context"
	"errors"
	"log"
	"mymod/api/auth"
	pb "mymod/genproto/authentication"
)



func (a *UserService) Register(ctx context.Context, req *pb.UserDetails) (*pb.ID, error) {
	id, err := a.Repo.Register(ctx, req)
    if err!= nil {
        return nil, err
    }
    return id, nil
	
}

func (a *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	resp, err := a.Repo.GetByUsername(ctx, req.Username)
	if err!= nil {
        return nil, errors.New("Invalid login")
    }
	if resp.Password!= req.Password {
        return nil, errors.New("Invalid password")
    }

	respose := &pb.LoginResponse{
		Access: &pb.AccessToken{
			Id:       resp.Id,
			Username: resp.Username,
			Email:    resp.Email,
		},
		Refresh: &pb.RefreshToken{
			Userid:     resp.Id,
		},
	}

	err = auth.GeneratedRefreshJWTToken(respose)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	err = auth.GeneratedAccessJWTToken(respose)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	err = a.Repo.StoreRefreshToken(ctx, respose)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return respose, nil

}

func (a *UserService) RefreshToken(ctx context.Context, req *pb.CheckRefreshTokenRequest) (*pb.CheckRefreshTokenResponse, error) {
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

	user, err := a.Repo.GetUserByID(ctx, id)
	if err!= nil {
        return nil, errors.New("Invalid refresh token")
    }

	res := pb.LoginResponse{
		Access: &pb.AccessToken{
			Id:       user.Id,
            Username: user.Username,
            Email:    user.Email,
		},
	}

	err = auth.GeneratedAccessJWTToken(&res)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &pb.CheckRefreshTokenResponse{Acces: true, Accestoken: res.Access.Accesstoken}, nil
}