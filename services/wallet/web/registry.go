package web

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/scalent-io/orchestration-framework/internal/middleware"
	zLog "github.com/scalent-io/orchestration-framework/pkg/log"
	"github.com/scalent-io/orchestration-framework/pkg/server"
	walletService "github.com/scalent-io/orchestration-framework/services/wallet/service"
)

type UserHandlerRegistryOptions struct {
	Config        *server.Config
	Logg          *zLog.LoggerConfig
	Middleware    middleware.Middleware
	WalletService walletService.WalletService
}

type UserHandlerRegistry struct {
	options UserHandlerRegistryOptions
}

func NewUserHandlerRegistry(options UserHandlerRegistryOptions) *UserHandlerRegistry {
	return &UserHandlerRegistry{options: options}
}

func (h *UserHandlerRegistry) StartServer() error {

	router, err := h.registerRoutes()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("server started successfully")
	router.Run(h.options.Config.Port)

	return nil
}

func (h *UserHandlerRegistry) registerRoutes() (*gin.Engine, error) {
	r := gin.Default()
	r.Use(h.options.Middleware.Cors())

	r.GET("/ping", h.homeHandler)

	walletRouter := r.Group("/wallet")
	walletRouter.GET("/", h.GetWalletHandler)

	return r, nil
}
