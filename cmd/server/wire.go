//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/ZQCard/kbk-layout/internal/biz"
	"github.com/ZQCard/kbk-layout/internal/conf"
	"github.com/ZQCard/kbk-layout/internal/data"
	"github.com/ZQCard/kbk-layout/internal/server"
	"github.com/ZQCard/kbk-layout/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Env, *conf.Server, *conf.Data, *conf.Bootstrap, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
