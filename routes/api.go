package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tiketin-management-api-with-go/controller"
	"github.com/tiketin-management-api-with-go/middleware"
)

func userRoutes(ge *gin.Engine) {
	user := ge.Group("/api/users")
	{
		user.POST("/register", controller.RegisterUserHandle)
		user.POST("/login", controller.LoginUserHandle)
		user.GET("/me", middleware.CheckJwt(), controller.GetUserHandle)
	}
}

func InitRoute(ge *gin.Engine) {
	userRoutes(ge)
}
