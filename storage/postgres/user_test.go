package postgres

import (
	"context"
	"reflect"
	"testing"
	pb "user_service/genproto/user"
)

func TestGetUserById(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := NewUserRepository(db)
	test := pb.ID{
		Id: "550e8400-e29b-41d4-a716-446655440000",
	}
	user, err := repo.GetUserByID(context.Background(), test.Id)
	if err != nil {
		t.Fatal(err)
	}

	exp := &pb.UserInfo{
		Id:       "550e8400-e29b-41d4-a716-446655440000",
		Username: "john_doe",
		Email:    "john.doe@example.com",
		Password: "password123",
	}

	if !reflect.DeepEqual(user, exp) {
		t.Error("Expected:", exp, "got:", user)
	}
}

func TestUpdateUser(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := NewUserRepository(db)
	test := pb.UserInfo{
		Id:       "123e4567-e89b-12d3-a456-426614174000",
		Username: "jane_smith",
		Email:    "jane.smith@example.com",
		Password: "securepass",
	}

	err = repo.UpdateUser(context.Background(), &test)
	if err != nil {
		t.Fatal(err)
	}

	user, err := repo.GetUserByID(context.Background(), test.Id)
	if err != nil {
		t.Fatal(err)
	}

	exp := &pb.UserInfo{
		Id:       "123e4567-e89b-12d3-a456-426614174000",
		Username: "jane_smith",
		Email:    "jane.smith@example.com",
		Password: "securepass",
	}

	if !reflect.DeepEqual(user, exp) {
		t.Error("Expected:", exp, "got:", user)
	}
}

func TestDeleteUser(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	test := pb.ID{
		Id: "456e8400-e29b-41d4-a716-446655440000",
	}

	err = repo.DeleteUser(context.Background(), test.Id)
	if err != nil {
		t.Fatal(err)
	}

	if err != nil {
		t.Error("User not found after deletion")
	}
}
