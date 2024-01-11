package middleware

import (
	"elementary-admin/constants"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"strings"
	"time"
)

func AuthenticationJwt() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			errMsg := fmt.Errorf("token doesn't Found")
			return c.Status(constants.CodeForbidden).JSON(constants.JsonRes(constants.CodeForbidden, errMsg))
		}

		tokenString := ExtractToken(c)
		err := ValidateToken(tokenString)

		if err != nil {
			errMsg := fmt.Errorf("token is not Valid")
			return c.Status(constants.CodeForbidden).JSON(constants.JsonRes(constants.CodeForbidden, errMsg))
		}

		return c.Next()
	}
}

func ExtractToken(c *fiber.Ctx) string {
	bearerToken := c.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ValidateToken(signToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}

	return
}

type JwtClaim struct {
	Username string `json:"username"`
	IdAdmin  int64  `json:"id_admin"`
	Position string `json:"position"`

	jwt.StandardClaims
}

var jwtKey = []byte("@@DMP170797XXS3CR3TXX@@")
