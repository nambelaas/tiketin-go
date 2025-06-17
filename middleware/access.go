package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiketin-management-api-with-go/helper"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		claim, err := helper.GetJwtData(c)
		if err != nil {
			helper.PrintErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}

		if claim.Role != "admin" {
			helper.PrintErrorResponse(c, http.StatusForbidden, "hanya bisa diakses oleh admin")
			return
		}

		c.Next()
	}
}

func AdminOrOrganizerOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		claim, err := helper.GetJwtData(c)
		if err != nil {
			helper.PrintErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}

		if claim.Role != "admin" && claim.Role != "organizer" {
			helper.PrintErrorResponse(c, http.StatusForbidden, "hanya bisa diakses oleh admin atau organizer")
			return
		}

		c.Next()
	}
}

func OrganizerOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		claim, err := helper.GetJwtData(c)
		if err != nil {
			helper.PrintErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}

		if claim.Role != "organizer" {
			helper.PrintErrorResponse(c, http.StatusForbidden, "hanya bisa diakses oleh organizer")
			return
		}

		c.Next()
	}
}

func UserOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		claim, err := helper.GetJwtData(c)
		if err != nil {
			helper.PrintErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}

		if claim.Role != "user" {
			helper.PrintErrorResponse(c, http.StatusForbidden, "hanya bisa diakses oleh user")
			return
		}

		c.Next()
	}
}
