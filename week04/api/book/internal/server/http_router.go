package server

import (
	"github.com/gin-gonic/gin"
	"week04/app/book/internal/biz"
	"week04/app/book/internal/conf"
)

// NewHttpRouter set gin.Engine as http.Handler
func NewHttpRouter(options conf.Options, biz *biz.HTTPBookBusiness) *gin.Engine {
	gin.SetMode(options.Mode)
	router := gin.Default()
	// set book router
	bookRouter(router, biz)
	return router
}

func bookRouter(app *gin.Engine, biz *biz.HTTPBookBusiness) {
	v := app.Group("v1/books")

	v.GET("/", biz.ListBooks)
	v.POST("/", biz.CreateBook)
	v.GET("/:id", biz.QueryBookById)
	v.PUT("/:id", biz.UpdateBookById)
	v.DELETE("/:id", biz.DeleteBookById)
}