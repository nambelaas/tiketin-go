package helper

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"math/big"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/tiketin-management-api-with-go/structs"
)

func EncodeError(message string) error {
	encode, _ := json.Marshal(map[string]string{
		"Message": message,
	})

	return errors.New(string(encode))
}

func ValidationCheck(err error) error {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		var listError []error
		for _, fe := range ve {
			msg := fe.Field() + " failed on the '" + fe.Tag() + "' tag"
			listError = append(listError, errors.New(msg))
		}

		return errors.Join(listError...)
	}

	return nil
}

func GetJwtData(ctx *gin.Context) (result structs.ClaimJwt, err error) {
	auth, exists := ctx.Get("auth")
	if !exists {
		return result, errors.New("data authorization tidak ditemukan")
	}

	dataJwt, ok := auth.(*structs.ClaimJwt)
	if !ok {
		return result, errors.New("data authorization tidak valid")
	}

	result = *dataJwt

	return result, nil
}

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result[i] = charset[num.Int64()]
	}

	return string(result)
}
