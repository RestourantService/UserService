package auth

import (
	"log"
	pb "mymod/genproto/authentication"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	signingkey = "signingkey"
)

func GeneratedAccessJWTToken(req *pb.LoginResponse) error {

	token := *jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = req.Access.Id
	claims["ureqame"] = req.Access.Username
	claims["email"] = req.Access.Email
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	newToken, err := token.SignedString([]byte(signingkey))
	if err!= nil {
        log.Print(err)
        return err
    }

	req.Access.Accesstoken = newToken

	return nil
}

func ValidateAccessToken(tokenString string) (bool, error) {
	_, err := ExractAccessToken(tokenString)
	if err!= nil {
        return false, err
    }
	return true, nil
}

func ExractAccessToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {	
		return []byte(signingkey),nil
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



