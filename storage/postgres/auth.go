package postgres

import (
	"context"
	"time"
	pb "user_service/genproto/authentication"
)

func (u *UserRepo) Register(ctx context.Context, user *pb.UserDetails) (*pb.ID, error) {
	var id pb.ID
	query := `
	insert into users (
		username, email, password
	)
	values (
		$1, $2, $3
	)
	returning id`

	err := u.DB.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(&id.Id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (u *UserRepo) GetUserByUsername(ctx context.Context, username string) (*pb.UserInfo, error) {
	user := pb.UserInfo{Username: username}
	query := `
	select
		id, email, password
	from
		users
	where
		deleted_at is null and username = $1
	`

	err := u.DB.QueryRowContext(ctx, query, username).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) StoreRefreshToken(ctx context.Context, token *pb.TokenRequest) error {
	query := `
	insert into refresh_tokens (
		user_id, token, expires_at
	)
	values (
		$1, $2, $3
	)`

	_, err := u.DB.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	_, err = u.DB.Exec(query, token.UserId, token.Token, token.ExpiresAt)
	return err
}

func (u *UserRepo) ValidateRefreshToken(ctx context.Context, token string) (string, error) {
	query := `
	select
		user_id, expires_at
	from
		refresh_tokens
	where
		token = $1
	`

	var userID string
	var expiresAt int64
	err := u.DB.QueryRowContext(ctx, query, token).Scan(&userID, &expiresAt)
	if err != nil {
		return "", err
	}
	if time.Now().After(time.Unix(expiresAt, 0)) {
		return "", err
	}
	return userID, nil
}

func (u *UserRepo) DeleteRefreshToken(ctx context.Context, userID string) error {
	query := `
	delete from
		refresh_tokens
	where
		user_id = $1`

	_, err := u.DB.ExecContext(ctx, query, userID)
	return err
}
