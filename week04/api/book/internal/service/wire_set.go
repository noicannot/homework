package service

import (
	"week04/api/book/internal/biz"
	"week04/api/book/internal/data/ent"

	"github.com/google/wire"
)

// ProvideSet for service package ...
var ProvideSet = wire.NewSet(NewHTTPBookRepo, NewGRPCBookRepo, NewGRPCBookService)

func NewHTTPBookRepo(client *ent.Client) biz.HTTPBookRepo {
	return &HTTPBookService{Client: client}
}

func NewGRPCBookRepo(client *ent.Client) biz.GRPCBookRepo {
	return &GRPCBookService{Client: client}
}

func NewGRPCBookService(client *ent.Client) *GRPCBookService {
	return &GRPCBookService{Client: client}
}

//https://github.com/qmdx00/Go-Camp/blob/master/week04/app/book/internal/service/wire_set.go
