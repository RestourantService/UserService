package service

import (
	"context"
	"database/sql"
	"log/slog"
	pbr "user_service/genproto/reservation"
	pb "user_service/genproto/user"
	l "user_service/pkg/logger"
	"user_service/storage/postgres"

	"github.com/pkg/errors"
)

type UserService struct {
	pb.UnimplementedUserServer
	Repo              *postgres.UserRepo
	Log               *slog.Logger
	ReservationClient pbr.ReservationClient
}

func NewUserService(db *sql.DB, reser pbr.ReservationClient) *UserService {
	return &UserService{
		Repo:              postgres.NewUserRepository(db),
		Log:               l.NewLogger(),
		ReservationClient: reser,
	}
}

func (u *UserService) GetUser(ctx context.Context, req *pb.ID) (*pb.UserInfo, error) {
	resp, err := u.Repo.GetUserByID(ctx, req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read user")
	}

	return resp, nil
}

func (u *UserService) UpdateUser(ctx context.Context, req *pb.UserInfo) (*pb.Void, error) {
	err := u.Repo.UpdateUser(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update user")
	}

	return &pb.Void{}, nil
}

func (u *UserService) DeleteUser(ctx context.Context, req *pb.ID) (*pb.Void, error) {
	err := u.Repo.DeleteUser(ctx, req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to delete user")
	}

	status, err := u.ReservationClient.DeleteReservationByUserID(ctx, &pbr.ID{Id: req.Id})
	if err != nil {
		return nil, errors.Wrap(err, "failed to delete user reservations")
	}

	if !status.Successful {
		return nil, errors.New("deletion of user reservations unsuccessful")
	}

	return &pb.Void{}, nil
}

func (u *UserService) ValidateUser(ctx context.Context, req *pb.ID) (*pb.Status, error) {
	resp, err := u.Repo.ValidateUser(ctx, req.Id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to validate user")
	}

	return resp, nil
}
