package server

import "github.com/google/wire"

// ProvideSet for server package ...
var ProvideSet = wire.NewSet(NewGRPCServer, NewHttpServer, NewHttpRouter)