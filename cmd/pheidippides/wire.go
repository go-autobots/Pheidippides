/**
 * @project Pheidippides
 * @Author 27
 * @Description //TODO
 * @Date 2021/12/18 11:53 12月
 **/
package pheidippides

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"pheidippides/internal/biz"
	"pheidippides/internal/conf"
	"pheidippides/internal/data"
	"pheidippides/internal/server"
	"pheidippides/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Redis) (*kratos.App, error) {
	panic(wire.Build(server.SrvSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}



