package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/aniket0951/lib_managment/sqlc/db/api/dto"
	"github.com/aniket0951/lib_managment/sqlc/db/api/services"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	AddNewUser(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	PurchaseBookForUser(ctx *gin.Context)
	ReturnUserBook(ctx *gin.Context)
	GetUserWithBooks(ctx *gin.Context)
}

type userController struct {
	userService services.UserService
}

func NewUserController(service services.UserService) UserController {
	return &userController{
		userService: service,
	}
}

func (cont *userController) AddNewUser(ctx *gin.Context) {
	var req dto.CreateUserParamsDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadGateway, errResponse(err))
		return
	}

	user, err := cont.userService.AddUser(req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "User has been added", "user_data": user})
}

func (cont *userController) GetUser(ctx *gin.Context) {
	var req dto.GetUserRequestDTO
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	user, err := cont.userService.GetUser(req.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "User Fetch Success", "user_data": user})
}

func (cont *userController) GetUsers(ctx *gin.Context) {
	var req dto.ListUsersDTO

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	users, err := cont.userService.GetUsers(req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "User Fetch Success", "users": users})
}

func (cont *userController) UpdateUser(ctx *gin.Context) {
	var req dto.UpdateUserDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	user, err := cont.userService.UpdateUser(req)

	if err != nil {
		fmt.Println("Error : ", err)
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "user info updated successfully", "user_data": user})
}

func (cont *userController) DeleteUser(ctx *gin.Context) {
	var req dto.GetUserRequestDTO

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	err := cont.userService.DeleteUser(req.Id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "user deleted successfully"})
}

func (cont *userController) PurchaseBookForUser(ctx *gin.Context) {
	var req dto.PurchaseBookDTO

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	err := cont.userService.PurchaseBookForUser(req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "book has been purchase"})
}

func (cont *userController) ReturnUserBook(ctx *gin.Context) {
	var req dto.PurchaseBookDTO

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	err := cont.userService.ReturnUserBook(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "Book has been return to library"})
}

func (cont *userController) GetUserWithBooks(ctx *gin.Context) {
	var req dto.GetUserRequestDTO
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	user_book_data,err := cont.userService.GetUserWithBooks(req.Id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg":"data fetched success", "user_data": user_book_data})
}