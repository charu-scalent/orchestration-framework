//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/scalent-io/orchestration-framework/internal/middleware"
	"github.com/scalent-io/orchestration-framework/pkg/db/sqlx"

	"github.com/scalent-io/orchestration-framework/services/wallet/repo"
	"github.com/scalent-io/orchestration-framework/services/wallet/service"
	"github.com/scalent-io/orchestration-framework/services/wallet/web"
)

var UserModuleSet = wire.NewSet(
	wire.FieldsOf(new(*WalletConfig), "server", "db", "logger", "middleware"),
	sqlx.NewSqlDB,

	wire.Struct(new(web.UserHandlerRegistryOptions), "*"),
	web.NewUserHandlerRegistry,

	middleware.NewMiddlewareImpl,
	wire.Bind(new(middleware.Middleware), new(*middleware.MiddlewareImpl)),

	repo.NewWalletRepoImpl,
	wire.Bind(new(service.WalletRepo), new(*repo.WalletRepoImpl)),

	service.NewWalletServiceImpl,
	wire.Bind(new(service.WalletService), new(*service.WalletServiceImpl)),
)

func initServer(config *WalletConfig) (*web.UserHandlerRegistry, error) {
	panic(wire.Build(UserModuleSet))
}
