package services

import (
	"context"
	"fmt"
	"time"

	"github.com/aniket0951/lib_managment/sqlc/db/api/dto"
	db "github.com/aniket0951/lib_managment/sqlc/db/sqlc"
)

type BookService interface {
	AddBook(req dto.AddBookParamDTO) (dto.GetBookDTO, error)
	GetBookWithAuthor(bookId int) (interface{}, error)
	GetAvailableBooks(req dto.ListAuthorDTO) (interface{}, error)
	PurchaseBook(bookId int64) error
}

type bookService struct {
	store *db.Store
}

func NewBookService(db *db.Store) BookService {
	return &bookService{
		store: db,
	}
}

func (bookServ *bookService) AddBook(req dto.AddBookParamDTO) (dto.GetBookDTO, error) {
	args := db.CreateBookParams{
		BookName:    req.BookName,
		BookDesc:    req.BookDesc,
		AuthorID:    req.AuthorID,
		PublishDate: time.Now(),
	}

	book, err := bookServ.store.CreateBook(context.Background(), args)

	bookResult := dto.GetBookDTO{
		ID:          book.ID,
		BookName:    book.BookName,
		BookDesc:    book.BookDesc,
		AuthorID:    book.AuthorID,
		PublishDate: book.PublishDate,
		CreatedAt:   book.CreatedAt,
	}

	return bookResult, err
}

func (bookServ *bookService) GetBookWithAuthor(bookId int) (interface{}, error) {
	book, err := bookServ.store.BookWithAuthor(context.Background(), int64(bookId))

	author := dto.GetAuthor{
		ID:            book.AuthorID,
		AuthorName:    book.AuthorName,
		AuthorAddress: book.AuthorAddress,
		CreatedAt:     book.CreatedAt_2,
	}
	bookWithAuthor := dto.GetBookWithAuthorDTO{
		ID:          book.ID,
		BookName:    book.BookName,
		BookDesc:    book.BookDesc,
		AuthorID:    book.AuthorID,
		PublishDate: book.PublishDate,
		CreatedAt:   book.CreatedAt,
		Author:      author,
	}

	return bookWithAuthor, err
}

func (bookServ *bookService) GetAvailableBooks(req dto.ListAuthorDTO) (interface{}, error) {

	books, err := bookServ.store.AvailableBooks(context.Background())
	fmt.Println("Books : ", len(books))
	return books, err
}

func (bookServ *bookService) PurchaseBook(bookId int64) error {
	_, err := bookServ.store.PurchaseBook(context.Background(), bookId)
	return err
}
