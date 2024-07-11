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
		c.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
	}
	res, err := h.Auth.Login(c, &req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error2": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  res.Access.Accesstoken,
		"refresh_token": res.Refresh.Refreshtoken,
	})

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
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
	if !res.Acces {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Refresh token is invalid"})
	}

	c.JSON(http.StatusOK, gin.H{
		"accestoken": res.Accestoken,
	})
}
