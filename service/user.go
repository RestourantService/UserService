package service

import (
	"context"
	"database/sql"
	"log/slog"
	pbr "user_service/genproto/reservation"
	pb "user_service/genproto/user"
	"user_service/pkg/logger"
	"user_service/storage/postgres"

	"github.com/pkg/errors"
)

type UserService struct {
	pb.UnimplementedUserServer
	Repo              *postgres.UserRepo
	ReservationClient pbr.ReservationClient
	Logger            *slog.Logger
}

func NewUserService(db *sql.DB, reser pbr.ReservationClient) *UserService {
	return &UserService{
		Repo:              postgres.NewUserRepository(db),
		ReservationClient: reser,
		Logger:            logger.NewLogger(),
	}
}

func (u *UserService) GetUser(ctx context.Context, req *pb.ID) (*pb.UserInfo, error) {
	u.Logger.Info("Get user Method is starting")

	resp, err := u.Repo.GetUserByID(ctx, req.Id)
	if err != nil {
		u.Logger.Error(errors.Wrap(err, "failed to read user").Error())
		return nil, errors.Wrap(err, "failed to read user")

	}

	u.Logger.Info("GetUser has successfully finished")
	return resp, nil
}

func (u *UserService) UpdateUser(ctx context.Context, req *pb.UserInfo) (*pb.Void, error) {
	u.Logger.Info("Update user Method is starting")

	err := u.Repo.UpdateUser(ctx, req)
	if err != nil {
		u.Logger.Error(errors.Wrap(err, "failed to update user").Error())
		return nil, errors.Wrap(err, "failed to update user")
	}

	u.Logger.Info("UpdateUser has successfully finished")
	return &pb.Void{}, nil
}

func (u *UserService) DeleteUser(ctx context.Context, req *pb.ID) (*pb.Void, error) {
	u.Logger.Info("Delete user Method is starting")

	err := u.Repo.DeleteUser(ctx, req.Id)
	if err != nil {
		u.Logger.Error(errors.Wrap(err, "failed to delete user").Error())
		return nil, errors.Wrap(err, "failed to delete user")
	}

	status, err := u.ReservationClient.DeleteReservationByUserID(ctx, &pbr.ID{Id: req.Id})
	if err != nil {
		u.Logger.Error(errors.Wrap(err, "failed to delete user reservations").Error())
		return nil, errors.Wrap(err, "failed to delete user reservations")
	}

	if !status.Successful {
		u.Logger.Error("deletion of user reservations unsuccessful")
		return nil, errors.New("deletion of user reservations unsuccessful")
	}

	u.Logger.Info("DeleteUser has successfully finished")
	return &pb.Void{}, nil
}

func (u *UserService) ValidateUser(ctx context.Context, req *pb.ID) (*pb.Status, error) {
	u.Logger.Info("Validate user Method is starting")

	resp, err := u.Repo.ValidateUser(ctx, req.Id)
	if err != nil {
		u.Logger.Error(errors.Wrap(err, "failed to validate user").Error())
		return &pb.Status{Successful: false}, errors.Wrap(err, "failed to validate user")
	}

	u.Logger.Info("ValidateUser has successfully finished")
	return &pb.Status{Successful: resp}, nil
}
