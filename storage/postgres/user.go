package postgres

import (
	"context"
	"database/sql"
	"log"
	"time"
	pb "user_service/genproto/user"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) GetUserByID(ctx context.Context, id string) (*pb.UserInfo, error) {
	user := &pb.UserInfo{Id: id}
	query := `
	SELECT
		username, email, password
	from
		users
	where
		id = $1 and deleted_at IS NULL
	`
	row := r.DB.QueryRowContext(ctx, query, id)

	err := row.Scan(&user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("user not found")
			return nil, err
		}
		log.Println("failed to scan user")
		return nil, err
	}

	return user, nil
}

func (r *UserRepo) UpdateUser(ctx context.Context, user *pb.UserInfo) error {
	query := `
	UPDATE 
		users
	SET
		username=$1, email=$2, password=$3, updated_at=$4
	WHERE
		id=$5 AND deleted_at IS NULL`

	result, err := r.DB.ExecContext(ctx, query, user.Username, user.Email, user.Password, time.Now(), user.Id)
	if err != nil {
		log.Println("failed to update user")
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		log.Println("failed to check rows affected")
		return err
	}

	if count == 0 {
		log.Println("user not found")
		return sql.ErrNoRows
	}

	return nil
}

func (r *UserRepo) DeleteUser(ctx context.Context, id string) error {
	query := `
	UPDATE
		users 
	SET
		deleted_at=NOW() 
	WHERE
		deleted_at is null and id=$1`

	_, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		log.Println("failed to delete user")
		return err
	}

	return nil
}

func (r *UserRepo) ValidateUser(ctx context.Context, id string) (*pb.Status, error) {
	query := `
    select EXISTS (
		select 1
		from users
		where id = $1 AND deleted_at IS NULL
	)`

	var status bool
	err := r.DB.QueryRowContext(ctx, query, id).Scan(&status)
	if err != nil {
		log.Println("failed to scan user")
		return nil, err
	}

	return &pb.Status{Successful: status}, nil
}
