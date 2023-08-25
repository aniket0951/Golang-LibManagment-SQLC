package dto

import "time"

type CreateAuthorDTO struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required,min=1"`
}

type GetAuthorByIDDTO struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type ListAuthorDTO struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

type AuthorAddressDTO struct {
	ID             int64     `json:"id"`
	AddressLineOne string    `json:"address_line_one"`
	City           string    `json:"city"`
	State          string    `json:"state"`
	Country        string    `json:"country"`
	CreatedAt      time.Time `json:"created_at"`
}

type GetAuthor struct {
	ID                  int64             `json:"id"`
	AuthorName          string            `json:"author_name"`
	AuthorAddress       string            `json:"author_address"`
	CreatedAt           time.Time         `json:"created_at"`
	AuthorDetailAddress *AuthorAddressDTO `json:"author_address_detail,omitempty"`
}

type GetAuthorWithBooks struct {
	ID            int64        `json:"id"`
	AuthorName    string       `json:"author_name"`
	AuthorAddress string       `json:"author_address"`
	CreatedAt     time.Time    `json:"created_at"`
	Books         []GetBookDTO `json:"books"`
}

type AddBookParamDTO struct {
	BookName string `json:"book_name"`
	BookDesc string `json:"book_desc"`
	AuthorID int64  `json:"author_id"`
}

type GetBookDTO struct {
	ID            int64                 `json:"id"`
	BookName      string                `json:"book_name"`
	BookDesc      string                `json:"book_desc"`
	AuthorID      int64                 `json:"author_id"`
	PublishDate   time.Time             `json:"publish_date"`
	CreatedAt     time.Time             `json:"created_at"`
	BookManagment []GetBookManagmentDTO `json:"book_managment,omitempty"`
	Author        *GetAuthor            `json:"author,omitempty"`
}

type GetBookBYID struct {
	ID int64 `uri:"id" binding="required"`
}

type GetBookWithAuthorDTO struct {
	ID          int64     `json:"id"`
	BookName    string    `json:"book_name"`
	BookDesc    string    `json:"book_desc"`
	AuthorID    int64     `json:"author_id"`
	PublishDate time.Time `json:"publish_date"`
	CreatedAt   time.Time `json:"created_at"`
	Author      GetAuthor `json:"author"`
}

type GetBookManagmentDTO struct {
	ID            int64     `json:"id"`
	BookID        int64     `json:"book_id"`
	TotalQuantity int32     `json:"total_quantity"`
	TotalInLab    int32     `json:"total_in_lab"`
	TotalOutLab   int32     `json:"total_out_lab"`
	CreatedAt     time.Time `json:"created_at"`
}

type GetAuthorWithBooksManagment struct {
	ID            int64        `json:"id"`
	AuthorName    string       `json:"author_name"`
	AuthorAddress string       `json:"author_address"`
	CreatedAt     time.Time    `json:"created_at"`
	Books         []GetBookDTO `json:"books"`
}
