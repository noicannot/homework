package biz

import (
	"github.com/gin-gonic/gin"
)

type HTTPBookRepo interface {
	ListBooks(ctx *gin.Context)
	CreateBook(ctx *gin.Context)
	GetBookById(ctx *gin.Context)
	UpdateBookById(ctx *gin.Context)
	DeleteBookById(ctx *gin.Context)
}

type HTTPBookBusiness struct {
	repo HTTPBookRepo
}

// ListBooks ...
func (b *HTTPBookBusiness) ListBooks(ctx *gin.Context) {
	b.repo.ListBooks(ctx)
}

// CreateBook ...
func (b *HTTPBookBusiness) CreateBook(ctx *gin.Context) {
	b.repo.CreateBook(ctx)
}

// DeleteBookById ...
func (b *HTTPBookBusiness) DeleteBookById(ctx *gin.Context) {
	b.repo.DeleteBookById(ctx)
}

// QueryBookById ...
func (b *HTTPBookBusiness) QueryBookById(ctx *gin.Context) {
	b.repo.GetBookById(ctx)
}

// UpdateBookById ...
func (b *HTTPBookBusiness) UpdateBookById(ctx *gin.Context) {
	b.repo.UpdateBookById(ctx)
}