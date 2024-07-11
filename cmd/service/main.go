package main

import (
	"log"
	"net"
	"user_service/config"
	pb "user_service/genproto/user"
	"user_service/service"
	"user_service/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()
	lis, err := net.Listen("tcp", cfg.Server.USER_PORT)
	if err != nil {
		log.Fatalf("error while listening: %v", err)
	}
	defer lis.Close()

	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatalf("error while connecting to database: %v", err)
	}
	defer db.Close()

	userService := service.NewUserService(db)
	server := grpc.NewServer()
	pb.RegisterUserServer(server, userService)

	log.Printf("server listening at %v", lis.Addr())
	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("error while serving: %v", err)
	}
}
