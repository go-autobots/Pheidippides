// +build wireinject

/**
 * @project Pheidippides
 * @Author 27
 * @Description //TODO
 * @Date 2021/12/18 11:53 12月
 **/
package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"

	"Pheidippides/internal/conf"
	"Pheidippides/internal/server"
	"Pheidippides/internal/service"
)

// newApp init kratos application.
func initApp(*conf.Server, *conf.Redis) (*kratos.App, error) {
	panic(wire.Build(server.SrvSet, service.ProviderSet, newApp))
}
