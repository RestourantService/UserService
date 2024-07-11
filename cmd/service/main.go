package main

import (
	"fmt"
	"log"
	"net"
	"user_service/config"
	pb1 "user_service/genproto/authentication"
	pb "user_service/genproto/user"

	"user_service/service"
	"user_service/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatalf("error while connecting to database: %v", err)
	}
	defer db.Close()

	fmt.Println("Starting server...")
	cfg := config.Load()
	fmt.Println(cfg.Server.USER_PORT)
	lis, err := net.Listen("tcp", cfg.Server.USER_PORT)
	if err != nil {
		log.Fatalf("error while listening: %v", err)
	}
	defer lis.Close()

	userService := service.NewUserService(db)
	authService := service.NewAuthService(db)
	server := grpc.NewServer()
	pb.RegisterUserServer(server, userService)
	pb1.RegisterAuthenticationServer(server, authService)

	log.Printf("server listening at %v", lis.Addr())
	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("error while serving: %v", err)
	}
}
