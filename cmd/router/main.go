package main

import (
	"fmt"
	"log"
	"user_service/api"
	"user_service/api/handler"
	"user_service/genproto/authentication"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	hand := Newhandler()
	router := api.Router(hand)
	fmt.Println("ok")
	log.Fatal(router.Run(":8085"))
}
func Newhandler() *handler.Handler {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	authservice := authentication.NewAuthenticationClient(conn)

	return &handler.Handler{Auth: authservice}
}
