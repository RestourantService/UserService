package handler


import (
	"mymod/genproto/authentication"
)

type Handler struct {
	Auth authentication.AuthenticationClient
}