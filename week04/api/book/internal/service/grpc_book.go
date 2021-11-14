package service

import (
	"context"
	"week04/api/book/internal/biz"
	"week04/api/book/internal/data/ent"
	where "week04/api/book/internal/data/ent/book"
	v1 "week04/api/book/service/v1"
	// "week04/app/book/internal/biz"
	// // "week04/app/book/internal/data/ent"
	// "week04/app/book/internal/data/ent"
	// where "week04/app/book/internal/data/ent/book"
)

var _ biz.GRPCBookRepo = (*GRPCBookService)(nil)

type GRPCBookService struct {
	v1.UnimplementedBookServer

	Client *ent.Client
}

func (g GRPCBookService) ListBooks(ctx context.Context, req *v1.ListBooksReq) (*v1.ListBooksReply, error) {
	items := make([]*v1.Item, 0)
	books, err := g.Client.Book.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	for _, book := range books {
		items = append(items, &v1.Item{
			Name:   book.Name,
			Author: book.Author,
			Number: int64(book.Number),
		})
	}
	return &v1.ListBooksReply{Books: items}, nil
}

func (g GRPCBookService) CreateBook(ctx context.Context, req *v1.CreateBookReq) (*v1.CreateBookReply, error) {
	item := req.Book
	book, err := g.Client.Book.
		Create().
		SetName(item.GetName()).
		SetAuthor(item.GetAuthor()).
		SetNumber(int(item.GetNumber())).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.CreateBookReply{Book: &v1.Item{
		Name:   book.Name,
		Author: book.Author,
		Number: int64(book.Number),
	}}, nil
}

func (g GRPCBookService) UpdateBook(ctx context.Context, req *v1.UpdateBookReq) (*v1.UpdateBookReply, error) {
	if err := g.Client.Book.
		UpdateOneID(int(req.Id)).
		SetName(req.Book.GetName()).
		SetAuthor(req.Book.GetAuthor()).
		SetNumber(int(req.Book.GetNumber())).
		Exec(ctx); err != nil {
		return nil, err
	}
	return &v1.UpdateBookReply{Book: req.Book}, nil
}

func (g GRPCBookService) GetBook(ctx context.Context, req *v1.GetBookReq) (*v1.GetBookReply, error) {
	book, err := g.Client.Book.
		Query().
		Where(where.IDEQ(int(req.Id))).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.GetBookReply{Book: &v1.Item{
		Name:   book.Name,
		Author: book.Author,
		Number: int64(book.Number),
	}}, nil
}

func (g GRPCBookService) DeleteBook(ctx context.Context, req *v1.DeleteBookReq) (*v1.DeleteBookReply, error) {
	if err := g.Client.Book.
		DeleteOneID(int(req.Id)).
		Exec(ctx); err != nil {
		return nil, err
	}
	return &v1.DeleteBookReply{}, nil
}
