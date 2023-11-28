package web

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/scalent-io/orchestration-framework/internal/middleware"
	zLog "github.com/scalent-io/orchestration-framework/pkg/log"
	"github.com/scalent-io/orchestration-framework/pkg/server"
	"github.com/scalent-io/orchestration-framework/services/order/service"
)

type OrderHandlerRegistryOptions struct {
	Config       *server.Config
	Logg         *zLog.LoggerConfig
	OrderService service.OrderService
	Middleware   middleware.Middleware
}

type OrderHandlerRegistry struct {
	options OrderHandlerRegistryOptions
}

func NewHandlerRegistry(options OrderHandlerRegistryOptions) *OrderHandlerRegistry {
	return &OrderHandlerRegistry{
		options: options,
	}
}

func (h *OrderHandlerRegistry) StartServer() error {

	router, err := h.registerRoutes()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("server started successfully")
	router.Run(h.options.Config.Port)

	return nil
}

func (h *OrderHandlerRegistry) registerRoutes() (*gin.Engine, error) {
	r := gin.Default()
	r.Use(h.options.Middleware.Cors())

	orderRouter := r.Group("/order")
	orderRouter.GET("/", h.GetOrderHandler)

	return r, nil
}
