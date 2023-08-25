package controllers

import (
	"net/http"

	"github.com/aniket0951/lib_managment/sqlc/db/api/dto"
	"github.com/aniket0951/lib_managment/sqlc/db/api/services"
	"github.com/gin-gonic/gin"
)

type BookController interface {
	AddBook(ctx *gin.Context)
	GetBook(ctx *gin.Context)
	GetBooks(ctx *gin.Context)
	PurchaseBook(ctx *gin.Context)
}

type bookController struct {
	bookService services.BookService
}

func NewBookController(service services.BookService) BookController {
	return &bookController{
		bookService: service,
	}
}

func (cont *bookController) AddBook(ctx *gin.Context) {
	var req dto.AddBookParamDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	book, err := cont.bookService.AddBook(req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (cont *bookController) GetBook(ctx *gin.Context) {
	var req dto.GetBookBYID

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	book, err := cont.bookService.GetBookWithAuthor(int(req.ID))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (cont *bookController) GetBooks(ctx *gin.Context) {
	var req dto.ListAuthorDTO

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	books, err := cont.bookService.GetAvailableBooks(req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "Books Fetch", "books": books})
}

func (cont *bookController) PurchaseBook(ctx *gin.Context) {
	var req dto.GetBookBYID

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	err := cont.bookService.PurchaseBook(req.ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "Book has been purchase"})
}
