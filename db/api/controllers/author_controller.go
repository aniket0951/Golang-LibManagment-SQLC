package controllers

import (
	"database/sql"
	"net/http"

	"github.com/aniket0951/lib_managment/sqlc/db/api/dto"
	"github.com/aniket0951/lib_managment/sqlc/db/api/services"
	"github.com/gin-gonic/gin"
)

type AuthorController interface {
	CreateNewAuthor(ctx *gin.Context)
	GetAuthor(ctx *gin.Context)
	GetAuthors(ctx *gin.Context)
	GetAuthorWithBooks(ctx *gin.Context)
	GetAuthorWithBooksAndManagment(ctx *gin.Context)
}

type authorController struct {
	authorService services.AuthorService
}

func NewAuthorController(service services.AuthorService) AuthorController {
	return &authorController{
		authorService: service,
	}
}

func (cont *authorController) CreateNewAuthor(ctx *gin.Context) {
	var req dto.CreateAuthorDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	author, err := cont.authorService.CreateAuthor(req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, author)
}

func (cont *authorController) GetAuthor(ctx *gin.Context) {
	var req dto.GetAuthorByIDDTO

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	author, err := cont.authorService.GetAuthor(int(req.ID))

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, author)
}

func (cont *authorController) GetAuthors(ctx *gin.Context) {
	var req dto.ListAuthorDTO

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	authors, err := cont.authorService.GetAuthors(req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, authors)
}

func (cont *authorController) GetAuthorWithBooks(ctx *gin.Context) {
	var authId dto.GetBookBYID
	if err := ctx.ShouldBindUri(&authId); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	author, err := cont.authorService.GetAuthorWithBooks(authId.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, author)
}

func (cont *authorController) GetAuthorWithBooksAndManagment(ctx *gin.Context) {
	var authId dto.GetBookBYID
	if err := ctx.ShouldBindUri(&authId); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}


	author, err := cont.authorService.GetAuthorWithBooksAndManagment(authId.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, author)
}

func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
