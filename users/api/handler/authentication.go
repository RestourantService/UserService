package handler

import (
	pb "mymod/genproto/authentication"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h *Handler)Register(c *gin.Context) {
	var req pb.UserDetails

	if err := c.ShouldBindJSON(&req); err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

	resp, err := h.Auth.Register(c, &req)
	if err!= nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) Login(c *gin.Context) {
	login := pb.LoginRequest{}
	if err := c.ShouldBindJSON(&login); err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

	resp, err := h.Auth.Login(c, &login)
	if err!= nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
	c.JSON(http.StatusOK, gin.H{
		"access_token":  resp.Access.Accesstoken,
        "refresh_token": resp.Refresh.Refreshtoken,
	})
}

func (h *Handler) RefreshToken(c *gin.Context) {
	RefreshToken := c.GetHeader("Authorization")

	if RefreshToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"Error": "unauthorization",
		})
		return
	}

	res, err := h.Auth.CheckRefreshToken(c, &pb.CheckRefreshTokenRequest{Token: RefreshToken})
	if err!= nil {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": err.Error()})
    }

	if !res.Acces {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": "refresh token is invalid"})
	}

	c.JSON(http.StatusOK,gin.H{"accesstoken": res.Accestoken})


}



