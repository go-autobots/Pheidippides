/**
 * @project Pheidippides
 * @Author 27
 * @Description //TODO
 * @Date 2021/12/18 11:53 12月
 **/
package main

import (
	"Pheidippides/internal/conf"

	"flag"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "/Users/fzk27/self-dev/Pheidippides/configs/config.yaml", "config path, eg: -conf config.yaml")
	flag.StringVar(&Name, "name", "Pheidippides", "App Name, eg: -name value")
	flag.StringVar(&Version, "version", "0.1", "App version, eg: -version value")
}

func newApp(gs *grpc.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Server(gs),
	)
}

func main() {
	flag.Parse()
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bootstrapConfig conf.Bootstrap
	if err := c.Scan(&bootstrapConfig); err != nil {
		panic(err)
	}

	app, err := initApp(bootstrapConfig.Server, bootstrapConfig.Redis)
	if err != nil {
		panic(err)
	}

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
