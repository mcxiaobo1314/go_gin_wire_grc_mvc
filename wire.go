//go:build wireinject
// +build wireinject

package main

import (
	"gin_gorm/config"
	"gin_gorm/controllers"
	"gin_gorm/core"
	"gin_gorm/middlewares"
	"gin_gorm/model"
	"gin_gorm/router"

	"github.com/google/wire"
)

func NewApp() *core.HServer {
	wire.Build(
		model.EnterModelProviderSet,
		config.EnterConfigProviderSet,
		controllers.ControllersProviderSet,
		router.RouterProviderSet,
		middlewares.EnterMiddlewareProviderSet,
		core.CoreProviderSet,
	)
	return nil
}
