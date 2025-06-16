package structs

import "github.com/golang-jwt/jwt/v5"

type ClaimJwt struct {
	jwt.RegisteredClaims
	UserId int
	Name   string
	Role   string
}
