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

func eventTypeRoutes(ge *gin.Engine) {
	eventType := ge.Group("/api/event-types")
	{
		eventType.POST("/create", middleware.CheckJwt(), middleware.AdminOnly(), controller.CreateEventTypeHandler)
		eventType.GET("/", controller.GetAllEventType)
		eventType.GET("/:id/events", controller.GetEventByEventTypeId)
		eventType.PUT("/:id", middleware.CheckJwt(), middleware.AdminOnly(), controller.UpdateEventTypeHandler)
		eventType.DELETE("/:id", middleware.CheckJwt(), middleware.AdminOnly(), controller.DeleteEventTypeHandler)
	}
}

func InitRoute(ge *gin.Engine) {
	userRoutes(ge)
	eventTypeRoutes(ge)
}
