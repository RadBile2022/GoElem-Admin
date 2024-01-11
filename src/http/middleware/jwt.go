package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateJwt(username string, idAdmin int64, position string) (tokenString string, err error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claim := &JwtClaim{
		Username: username,
		IdAdmin:  idAdmin,
		Position: position,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err = token.SignedString(jwtKey)
	return
}
