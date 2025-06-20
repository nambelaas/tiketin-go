package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
		eventType.POST("/create", middleware.CheckJwt(), middleware.AdminOnly(), controller.CreateEventTypeHandle)
		eventType.GET("/list", controller.GetAllEventTypeHandle)
		eventType.GET("/:id/events", controller.GetEventByEventTypeIdHandle)
		eventType.PUT("/:id/update", middleware.CheckJwt(), middleware.AdminOnly(), controller.UpdateEventTypeHandle)
		eventType.DELETE("/:id/delete", middleware.CheckJwt(), middleware.AdminOnly(), controller.DeleteEventTypeHandle)
	}
}

func eventRoutes(ge *gin.Engine) {
	eventType := ge.Group("/api/events")
	{
		eventType.POST("/create", middleware.CheckJwt(), middleware.AdminOrOrganizerOnly(), controller.CreateEventHandle)
		eventType.GET("/list", controller.GetAllEventHandle)
		eventType.GET("/:event_id", controller.GetEventByIdHandle)
		eventType.GET("/me", middleware.CheckJwt(), middleware.AdminOrOrganizerOnly(), controller.GetEventByUserHandle)
		eventType.PUT("/:event_id/update", middleware.CheckJwt(), middleware.AdminOrOrganizerOnly(), controller.UpdateEventHandle)
		eventType.DELETE("/:event_id/delete", middleware.CheckJwt(), middleware.AdminOrOrganizerOnly(), controller.DeleteEventHandle)
	}
}

func ticketRoutes(ge *gin.Engine) {
	eventType := ge.Group("/api/events/:event_id/tickets")
	{
		eventType.POST("/create", middleware.CheckJwt(), middleware.AdminOrOrganizerOnly(), controller.CreateTicketHandle)
		eventType.GET("/list", controller.GetAllTicketEventHandle)
		eventType.PUT("/:ticket_id/update", middleware.CheckJwt(), middleware.AdminOrOrganizerOnly(), controller.UpdateTicketHandle)
		eventType.DELETE("/:ticket_id/delete", middleware.CheckJwt(), middleware.AdminOrOrganizerOnly(), controller.DeleteTicketHandle)
	}
}

func reviewRoutes(ge *gin.Engine) {
	eventType := ge.Group("/api/events/:event_id/reviews")
	{
		eventType.POST("/create", middleware.CheckJwt(), middleware.UserOnly(), controller.CreateReviewHandle)
		eventType.GET("/list", controller.GetAllReviewEventHandle)
		eventType.PUT("/:review_id/update", middleware.CheckJwt(), middleware.UserOnly(), controller.UpdateReviewHandle)
	}
}

func orderRoutes(ge *gin.Engine) {
	eventType := ge.Group("/api/orders")
	{
		eventType.POST("/create", middleware.CheckJwt(), middleware.UserOnly(), controller.CreateOrderHandle)
		eventType.PUT("/:order_id/pay", middleware.CheckJwt(), middleware.UserOnly(), controller.PayOrderHandle)
		eventType.PUT("/:order_id/cancel", middleware.CheckJwt(), middleware.UserOnly(), controller.CancelOrderHandle)
		eventType.GET("/list", middleware.CheckJwt(), controller.GetAllOrderHandle)
		eventType.GET("/checkin/ticket", controller.CheckInHandle)
		eventType.GET("/user/me", middleware.CheckJwt(), middleware.UserOnly(), controller.GetOrderByUserAuthHandle)
		eventType.GET("/user/:user_id/list", middleware.CheckJwt(), middleware.AdminOrOrganizerOnly(), controller.GetOrderByUserHandle)
		eventType.GET("/:order_id", middleware.CheckJwt(), controller.GetOrderByIdHandle)
	}
}

func InitRoute(ge *gin.Engine) {
	userRoutes(ge)
	eventTypeRoutes(ge)
	eventRoutes(ge)
	ticketRoutes(ge)
	reviewRoutes(ge)
	orderRoutes(ge)

	ge.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
