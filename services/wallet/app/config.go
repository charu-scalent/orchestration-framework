package main

import (
	"github.com/scalent-io/orchestration-framework/internal/middleware"
	"github.com/scalent-io/orchestration-framework/pkg/db/sqlx"
	"github.com/scalent-io/orchestration-framework/pkg/log"
	"github.com/scalent-io/orchestration-framework/pkg/server"
	"github.com/spf13/viper"
)

type WalletConfig struct {
	Server     *server.Config
	DB         *sqlx.DbConfig
	Logger     *log.LoggerConfig
	Middleware middleware.MiddlewareConfig
}

var config WalletConfig

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("user")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	viper.Unmarshal(&config)
	return nil
}
