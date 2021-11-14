package biz

import (
	"context"
	v1 "week04/api/book/internal/service/v1"
)

type GRPCBookRepo interface {
	ListBooks(context.Context, *v1.ListBooksReq) (*v1.ListBooksReply, error)
	CreateBook(context.Context, *v1.CreateBookReq) (*v1.CreateBookReply, error)
	UpdateBook(context.Context, *v1.UpdateBookReq) (*v1.UpdateBookReply, error)
	GetBook(context.Context, *v1.GetBookReq) (*v1.GetBookReply, error)
	DeleteBook(context.Context, *v1.DeleteBookReq) (*v1.DeleteBookReply, error)
}

type GRPCBookBusiness struct {
	repo GRPCBookRepo
}
