package service

import (
	"context"
	"net/http"
	"week04/api/book/internal/biz"
	"week04/api/book/internal/data/ent"
	where "week04/api/book/internal/data/ent/book"
	"week04/api/book/internal/pkg"

	// "week04/app/book/internal/pkg"
	// "week04/api/book/internal/pkg"

	"github.com/gin-gonic/gin"
)

var _ biz.HTTPBookRepo = (*HTTPBookService)(nil)

type HTTPBookService struct {
	Client *ent.Client
}

func (s *HTTPBookService) ListBooks(ctx *gin.Context) {
	books, err := s.Client.Book.
		Query().
		All(context.Background())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, books)
}

func (s *HTTPBookService) CreateBook(ctx *gin.Context) {
	name := ctx.Query("name")
	author := ctx.Query("author")

	book, err := s.Client.Book.
		Create().
		SetName(name).
		SetAuthor(author).
		SetNumber(10).
		Save(context.Background())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, book)
}

func (s *HTTPBookService) GetBookById(ctx *gin.Context) {
	id := pkg.StringToInt(ctx.Param("id"))

	book, err := s.Client.
		Book.
		Query().
		Where(where.IDEQ(id)).
		Only(context.Background())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, book)
}

func (s *HTTPBookService) UpdateBookById(ctx *gin.Context) {
	id := pkg.StringToInt(ctx.Param("id"))
	name := ctx.Query("name")
	author := ctx.Query("author")
	number := pkg.StringToInt(ctx.Query("number"))

	if err := s.Client.Book.
		UpdateOneID(id).
		SetName(name).
		SetAuthor(author).
		SetNumber(number).
		Exec(context.Background()); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func (s *HTTPBookService) DeleteBookById(ctx *gin.Context) {
	id := pkg.StringToInt(ctx.Param("id"))

	if err := s.Client.Book.
		DeleteOneID(id).
		Exec(context.Background()); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
