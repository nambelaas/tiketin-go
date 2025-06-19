package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/skip2/go-qrcode"
	"github.com/spf13/viper"
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

func GenerateQRCode(orderId int, orderItemID int, ticketTypeId int) (string, error) {
	baseURL := viper.GetString("App.BaseUrl")
	qrContentURL := fmt.Sprintf("%s/api/orders/checkin/ticket?orderId=%d&orderItemId=%d&ticketId=%d", baseURL, orderId, orderItemID, ticketTypeId)

	filename := fmt.Sprintf("ticket_%05d_%05d.png", orderItemID, ticketTypeId)
	localPath := fmt.Sprintf("public/assets/qrcode/%s", filename)

	err := os.MkdirAll("public/assets/qrcode", os.ModePerm)
	if err != nil {
		return "", err
	}

	err = qrcode.WriteFile(qrContentURL, qrcode.Medium, 256, localPath)
	if err != nil {
		return "", err
	}

	publicURL := fmt.Sprintf("%s/assets/qrcode/%s", baseURL, filename)

	return publicURL, nil
}
