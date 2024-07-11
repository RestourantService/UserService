package api

import (
	"user_service/api/handler"

	"github.com/gin-gonic/gin"
)

func Router(hand *handler.Handler) *gin.Engine {
	router := gin.Default()
	user := router.Group("/user")
	user.POST("/register", hand.Register)
	user.POST("/login", hand.Login)
	user.GET("/refresh", hand.RefreshToken)
	return router
}
