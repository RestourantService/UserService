package handler

import (
	"fmt"
	"net/http"
	pb "user_service/genproto/authentication"

	"github.com/gin-gonic/gin"
)

func (h Handler) Register(c *gin.Context) {
	h.Log.Info("Register is starting")
	req := pb.UserDetails{}
	if err := c.BindJSON(&req); err != nil {
		h.Log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	res, err := h.Auth.Register(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		er := fmt.Sprintf("Error to Insert user : %s", err)
		c.Writer.Write([]byte(er))
		return
	}
	h.Log.Info("Register ended")
	c.JSON(http.StatusOK, res)
}

func (h Handler) Login(c *gin.Context) {
	h.Log.Info("Login is working")
	req := pb.LoginRequest{}
	if err := c.BindJSON(&req); err != nil {
		h.Log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
	}
	res, err := h.Auth.Login(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error2": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  res.Access.Accesstoken,
		"refresh_token": res.Refresh.Refreshtoken,
	})
	h.Log.Info("login is succesfully ended")

}

func (h Handler) RefreshToken(c *gin.Context) {
	h.Log.Info("refreshtoken is working")
	RefreshtokenStr := c.GetHeader("Authorization")

	if RefreshtokenStr == "" {
		h.Log.Error("Empty anauthorization")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}
	res, err := h.Auth.CheckRefreshToken(c, &pb.CheckRefreshTokenRequest{Token: RefreshtokenStr})
	if err != nil {
		h.Log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
	if !res.Acces {
		h.Log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Refresh token is invalid"})
	}

	c.JSON(http.StatusOK, gin.H{
		"accestoken": res.Accestoken,
	})
	h.Log.Info("refreshtoken is succesful")
}
