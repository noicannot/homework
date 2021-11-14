package biz

import "github.com/google/wire"

// ProvideSet for biz package ...
var ProvideSet = wire.NewSet(NewHttpBookBusiness)

func NewHttpBookBusiness(repo HTTPBookRepo) *HTTPBookBusiness {
	return &HTTPBookBusiness{repo: repo}
}

func NewGRPCBookBusiness(repo GRPCBookRepo) *GRPCBookBusiness {
	return &GRPCBookBusiness{repo: repo}
}