package handler

import "user_service/genproto/authentication"

type Handler struct {
	Auth authentication.AuthenticationClient
}
