package dto

import "time"

type CreateUserParamsDTO struct {
	UserName     string `json:"name" binding:"required"`
	UserEmail    string `json:"email" binding:"required,email"`
	UserPassword string `json:"password" binding:"required"`
}

type GetUserRequestDTO struct {
	Id int64 `uri:"id"`
}

type ListUsersDTO struct {
	PageID   int32 `form:"page_id" binding:"required"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

type UpdateUserDTO struct {
	Id       int64  `json:"id" binding:"required,min=1"`
	UserName string `json:"name"  binding:"required"`
}



type GetUsersDTO struct {
	ID           int64     `json:"id"`
	UserName     string    `json:"user_name"`
	UserEmail    string    `json:"user_email"`
	UserPassword string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	Books 		[]GetBookDTO `json:"books,omitempty"`
}

type PurchaseBookDTO struct {
	UserId int64 `form:"user_id" binding:"required"`
	BookId int64 `form:"book_id" binding:"required"`
}

