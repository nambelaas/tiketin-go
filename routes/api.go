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
		eventType.GET("/list", controller.GetAllEventType)
		eventType.GET("/:id/events", controller.GetEventByEventTypeId)
		eventType.PUT("/:id/update", middleware.CheckJwt(), middleware.AdminOnly(), controller.UpdateEventTypeHandler)
		eventType.DELETE("/:id/delete", middleware.CheckJwt(), middleware.AdminOnly(), controller.DeleteEventTypeHandler)
	}
}

func eventRoutes(ge *gin.Engine) {
	eventType := ge.Group("/api/events")
	{
		eventType.POST("/create", middleware.CheckJwt(), middleware.AdminOrOrganizerOnly(), controller.CreateEventHandler)
		eventType.GET("/list", controller.GetAllEvent)
		eventType.GET("/:event_id", controller.GetEventById)
		eventType.GET("/me", middleware.CheckJwt(), middleware.AdminOrOrganizerOnly(), controller.GetEventByUser)
		eventType.PUT("/:event_id/update", middleware.CheckJwt(), middleware.AdminOrOrganizerOnly(), controller.UpdateEventHandler)
		eventType.DELETE("/:event_id/delete", middleware.CheckJwt(), middleware.AdminOrOrganizerOnly(), controller.DeleteEventHandler)
	}
}

func InitRoute(ge *gin.Engine) {
	userRoutes(ge)
	eventTypeRoutes(ge)
	eventRoutes(ge)
}
