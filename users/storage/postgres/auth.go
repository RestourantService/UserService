package postgres

import (
	"context"
	"log"
	pb "mymod/genproto/authentication"
)



func (r *UserRepo) Register(ctx context.Context, user *pb.UserDetails) (*pb.ID, error) {
	query := `
		insert into auth (username, email, password)
		values ($1, $2, $3)
		returning id
		`

	var id pb.ID

	err := r.DB.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(&id.Id)
	if err != nil {
		log.Println("failed to insert auth")
		return nil, err
	}
	return &id, nil
}

func (r *UserRepo) GetByUsername(ctx context.Context, username string) (*pb.UserInfo, error) {
	query := `
		select id, username, email
		from auth
		where deleted_at IS NULL and username = $1
		`

	 resp := pb.UserInfo{Username: username}

	err := r.DB.QueryRowContext(ctx, query, username).Scan(&resp.Id, &resp.Username, &resp.Email)
	if err != nil {
		log.Println("user not found")
		return nil, err
	}
	
	return &resp, nil
}

func (r *UserRepo) StoreRefreshToken(ctx context.Context, token *pb.LoginResponse) error {
	query := `
	    insert into refresh_token (user_id, token)
		values ($1, $2)
		`
	
	_, err := r.DB.ExecContext(ctx, query, token.Refresh.Userid, token.Refresh.Refreshtoken)
	if err!= nil {
        log.Println("failed to insert refresh token")
        return err
    }
	return nil
}

func (r *UserRepo) DeleteRefreshToken(ctx context.Context, userId string) error {
	query := `
		delete from refresh_token
		where user_id = $1
		`
    _, err := r.DB.ExecContext(ctx, query, userId)
	if err!= nil {
        log.Println("failed to delete refresh token")
        return err
    }
	return nil
}
