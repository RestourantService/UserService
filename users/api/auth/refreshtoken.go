package auth

import (
	"log"
	pb "mymod/genproto/authentication"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	refreshTokenKey = "refreshTokenKey"
)

func GeneratedRefreshJWTToken(req *pb.LoginResponse) error {
	token := *jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = req.Access.Id
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	newToken, err := token.SignedString([]byte(refreshTokenKey))
	if err!= nil {
        log.Print(err)
        return err
    }

	req.Access.Accesstoken = newToken

	return nil
}

func ValidateRefreshToken(tokenString string) (bool, error) {
	_, err := ExtractRefreshToken(tokenString)
    if err!= nil {
        return false, err
    }
    return true, nil
}

func ExtractRefreshToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(refreshTokenKey), nil
	})
    if err!= nil {
        return nil, err
    }

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, err
	}

	return &claims, nil
}

func GetUserIdFromRefreshToken(refreshtokenString string) (string, error) {
	refreshtoken, err := jwt.Parse(refreshtokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(refreshTokenKey), nil
	})
	if err!= nil {
        return "", err
    }

	claims, ok := refreshtoken.Claims.(jwt.MapClaims)
	if !(ok && refreshtoken.Valid) {
		return "", err
	}

	UserId := claims["user_id"].(string)

	return UserId, nil
}