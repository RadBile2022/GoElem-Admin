package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func OnlyOwner(fh fiber.Handler) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ExtractToken(ctx)

		level, _ := ValidateTokenClaim(token)
		if level.Position == "owner" {
			return fh(ctx)
		} else {
			return ctx.SendStatus(http.StatusUnauthorized)
		}
	}
}

func ValidateTokenClaim(signToken string) (claim *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

	if err != nil {
		return
	}

	claim, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}

	return
}
