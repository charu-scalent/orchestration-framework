//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/scalent-io/orchestration-framework/internal/middleware"
	"github.com/scalent-io/orchestration-framework/pkg/db/sqlx"
	"github.com/scalent-io/orchestration-framework/services/order/repo"
	"github.com/scalent-io/orchestration-framework/services/order/service"
	"github.com/scalent-io/orchestration-framework/services/order/web"
)

var OrderModuleSet = wire.NewSet(

	wire.FieldsOf(new(*Config), "server", "db", "logger", "middleware"),
	sqlx.NewSqlDB,

	middleware.NewMiddlewareImpl,
	wire.Bind(new(middleware.Middleware), new(*middleware.MiddlewareImpl)),

	repo.NewOrderRepoImpl,
	wire.Bind(new(service.OrderRepo), new(*repo.OrderRepoImpl)),

	service.NewOrderServiceImpl,
	wire.Bind(new(service.OrderService), new(*service.OrderServiceImpl)),

	wire.Struct(new(web.OrderHandlerRegistryOptions), "*"),
	web.NewHandlerRegistry,
)

func initServer(config *Config) (*web.OrderHandlerRegistry, error) {
	panic(wire.Build(OrderModuleSet))
}
