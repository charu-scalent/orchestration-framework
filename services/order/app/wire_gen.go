// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/scalent-io/orchestration-framework/internal/middleware"
	"github.com/scalent-io/orchestration-framework/pkg/db/sqlx"
	"github.com/scalent-io/orchestration-framework/services/order/repo"
	"github.com/scalent-io/orchestration-framework/services/order/service"
	"github.com/scalent-io/orchestration-framework/services/order/web"
)

// Injectors from wire.go:

func initServer(config2 *Config) (*web.OrderHandlerRegistry, error) {
	serverConfig := config2.Server
	loggerConfig := config2.Logger
	dbConfig := config2.DB
	db, err := sqlx.NewSqlDB(dbConfig)
	if err != nil {
		return nil, err
	}
	orderRepoImpl, err := repo.NewOrderRepoImpl(db)
	if err != nil {
		return nil, err
	}
	orderServiceImpl, err := service.NewOrderServiceImpl(orderRepoImpl)
	if err != nil {
		return nil, err
	}
	middlewareImpl, err := middleware.NewMiddlewareImpl()
	if err != nil {
		return nil, err
	}
	orderHandlerRegistryOptions := web.OrderHandlerRegistryOptions{
		Config:       serverConfig,
		Logg:         loggerConfig,
		OrderService: orderServiceImpl,
		Middleware:   middlewareImpl,
	}
	orderHandlerRegistry := web.NewHandlerRegistry(orderHandlerRegistryOptions)
	return orderHandlerRegistry, nil
}

// wire.go:

var OrderModuleSet = wire.NewSet(wire.FieldsOf(new(*Config), "server", "db", "logger", "middleware"), sqlx.NewSqlDB, middleware.NewMiddlewareImpl, wire.Bind(new(middleware.Middleware), new(*middleware.MiddlewareImpl)), repo.NewOrderRepoImpl, wire.Bind(new(service.OrderRepo), new(*repo.OrderRepoImpl)), service.NewOrderServiceImpl, wire.Bind(new(service.OrderService), new(*service.OrderServiceImpl)), wire.Struct(new(web.OrderHandlerRegistryOptions), "*"), web.NewHandlerRegistry)
