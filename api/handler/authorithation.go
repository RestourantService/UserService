package handler

import (
	"fmt"
	"net/http"
	pb "user_service/genproto/authentication"

	"github.com/gin-gonic/gin"
)

func (h Handler) Register(c *gin.Context) {
	req := pb.UserDetails{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	res, err := h.Auth.Register(c, &req)
	if err != nil {
		er := fmt.Sprintf("Error to Insert user : %s", err)
		c.Writer.Write([]byte(er))
		return
	}
	c.JSON(http.StatusOK, res)

}

func (h Handler) Login(c *gin.Context) {
	req := pb.LoginRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	res, err := h.Auth.Login(c, &req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)

}

func (h Handler) RefreshToken(c *gin.Context) {
	RefreshtokenStr := c.GetHeader("Authorization")

	if RefreshtokenStr == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}
	res, err := h.Auth.CheckRefreshToken(c, &pb.CheckRefreshTokenRequest{Token: RefreshtokenStr})
	if !res.Acces {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Refresh token is invalid"})
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, res.Accestoken)
}
