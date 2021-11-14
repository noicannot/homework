// +build wireinject

package main

import (
	"github.com/google/wire"
	"golang.org/x/sync/errgroup"
	"week04/app/book/internal/biz"
	"week04/app/book/internal/conf"
	"week04/app/book/internal/data"
	"week04/app/book/internal/data/ent"
	"week04/app/book/internal/server"
	"week04/app/book/internal/service"
)

type App struct {
	HttpServer *server.HttpServer
	GRPCServer *server.GRPCServer
	Client     *ent.Client
}

// newApp return App struct with server.HttpServer and server.GRPCServer
func newApp(http *server.HttpServer, grpc *server.GRPCServer, client *ent.Client) *App {
	return &App{HttpServer: http, GRPCServer: grpc, Client: client}
}

// initApp Inject wire ProvideSet
func initApp(group *errgroup.Group, option conf.Options) *App {
	panic(wire.Build(server.ProvideSet, data.ProvideSet, service.ProvideSet, biz.ProvideSet, newApp))
}
