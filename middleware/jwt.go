package middleware

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"github.com/tiketin-management-api-with-go/helper"
	"github.com/tiketin-management-api-with-go/structs"
)

func GenerateJwtToken(u structs.Users) string {
	claims := structs.ClaimJwt{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    viper.GetString("JWT.Issuer"),
		},
		UserId: u.Id,
		Name: u.Name,
		Role: u.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	res, _ := token.SignedString([]byte(viper.GetString("JWT.SignatureKey")))

	return res
}

func GetJwtTokenFromHeader(c *gin.Context) (tokenString string, err error) {
	authHeader := c.Request.Header.Get("Authorization")

	if checkIsStringEmpty(authHeader) {
		return tokenString, errors.New("header authorization dibutuhkan")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return tokenString, errors.New("format header authorization tidak sesuai")
	}

	return parts[1], nil
}

func checkIsStringEmpty(input string) bool {
	return input == ""
}

func CheckJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := GetJwtTokenFromHeader(c)
		if err != nil {
			helper.PrintErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &structs.ClaimJwt{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("JWT.SignatureKey")), nil
		})

		if err != nil {
			helper.PrintErrorResponse(c, http.StatusBadRequest, "token tidak valid")
			return
		}

		claims, ok := token.Claims.(*structs.ClaimJwt)
		if !ok || !token.Valid {
			helper.PrintErrorResponse(c, http.StatusBadRequest, "token gagal diklaim")
			return
		}

		if claims.ExpiresAt == nil || claims.ExpiresAt.Time.Before(time.Now()) {
			helper.PrintErrorResponse(c, http.StatusBadRequest, "token kadaluarsa, silahkan login kembali")
			return
		}

		c.Set("auth", claims)
		c.Next()
	}
}
