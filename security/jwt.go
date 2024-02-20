package security

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kliffx2/trending-repo/model"
)

const SECRET_KEY = "kliffx2"

func GenToken(user model.User) (string, error){
	claims := &model.JwtCustomClaims{
		UserId:         user.UserId,
		Role:           user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}

	return result, nil
}