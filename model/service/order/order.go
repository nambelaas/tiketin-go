package order

import (
	"github.com/gin-gonic/gin"
	"github.com/tiketin-management-api-with-go/model/repository/order"
	"github.com/tiketin-management-api-with-go/structs"
)

type OrderServiceInterface interface {
	CreateOrder(ctx *gin.Context) error
	PayOrder(ctx *gin.Context) error
	CancelOrder(ctx *gin.Context) error
	GetAllOrder(ctx *gin.Context) ([]structs.Order, error)
	GetOrderByUser(ctx *gin.Context) ([]structs.Order, error)
	GetOrderByUserAuth(ctx *gin.Context) ([]structs.Order, error)
	GetOrderById(ctx *gin.Context) (structs.Order, error)
	CheckIn(ctx *gin.Context) error
}

type OrderService struct {
	repo order.OrderRepositoryInterface
}
