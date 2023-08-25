package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/aniket0951/lib_managment/sqlc/db/api/dto"
	db "github.com/aniket0951/lib_managment/sqlc/db/sqlc"
)

type AuthorService interface {
	CreateAuthor(dto.CreateAuthorDTO) (dto.GetAuthor, error)
	GetAuthor(int) (dto.GetAuthor, error)
	GetAuthors(req dto.ListAuthorDTO) ([]dto.GetAuthor, error)
	GetAuthorWithBooks(authorID int64) (interface{}, error)
	GetAuthorWithBooksAndManagment(authorID int64) (interface{}, error)
}

type authorService struct {
	store *db.Store
}

func NewAuthorService(store *db.Store) AuthorService {
	return &authorService{store: store}
}

func (authorSer *authorService) CreateAuthor(req dto.CreateAuthorDTO) (dto.GetAuthor, error) {

	args := db.CreateAuthorParams{
		AuthorName:    req.Name,
		AuthorAddress: req.Address,
	}

	author, err := authorSer.store.CreateAuthor(context.Background(), args)

	return dto.GetAuthor{
		ID: author.ID,
		AuthorName: author.AuthorName,
		AuthorAddress: author.AuthorAddress,
		CreatedAt: author.CreatedAt,
	}, err
}

func (authSer *authorService) GetAuthor(authorID int) (dto.GetAuthor, error) {

	author, err := authSer.store.GetAuthor(context.Background(), int64(authorID))

	return dto.GetAuthor{
		ID: author.ID,
		AuthorName: author.AuthorName,
		AuthorAddress: author.AuthorAddress,
		CreatedAt: author.CreatedAt,
	}, err
}

func (authServ *authorService) GetAuthors(req dto.ListAuthorDTO) ([]dto.GetAuthor, error) {

	args := db.GetAuthorsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	authors, err := authServ.store.GetAuthors(context.Background(), args)

	getAuthers := make([]dto.GetAuthor, len(authors))

	for i, author := range authors {
		getAuthers[i] = dto.GetAuthor{
			ID: author.ID,
			AuthorName: author.AuthorName,
			AuthorAddress: author.AuthorAddress,
			CreatedAt: author.CreatedAt,
		}
	}

	return getAuthers, err
}

func (authServ *authorService) GetAuthorWithBooks(authorID int64) (interface{}, error) {
	result, err := authServ.store.GetAuthorWithBooks(context.Background(), authorID)
	author := dto.GetAuthorWithBooks{}
	books := make([]dto.GetBookDTO, len(result))

	if len(result) > 0 {
		author.ID = result[0].ID
		author.AuthorName = result[0].AuthorName
		author.AuthorAddress = result[0].AuthorAddress
		author.CreatedAt = result[0].CreatedAt
	} else {
		return nil, errors.New(sql.ErrNoRows.Error())
	}

	for i, book := range result {
		tempBook := dto.GetBookDTO{
			ID:          book.ID_2,
			BookName:    book.BookName,
			BookDesc:    book.BookDesc,
			AuthorID:    book.AuthorID,
			PublishDate: book.PublishDate,
			CreatedAt:   book.CreatedAt,
		}

		books[i] = tempBook
	}

	author.Books = books

	return author, err
}

func (authServ *authorService) GetAuthorWithBooksAndManagment(authorID int64) (interface{}, error) {
	result, err := authServ.store.GetAuthorWithBooksAndManagment(context.Background(), authorID)

	if err != nil {
		return nil, err
	}
	author := dto.GetAuthorWithBooksManagment{}
	books := make([]dto.GetBookDTO, len(result))

	managemet := []dto.GetBookManagmentDTO{}

	if len(result) > 0 {
		author.ID = result[0].ID
		author.AuthorName = result[0].AuthorName
		author.AuthorAddress = result[0].AuthorAddress
		author.CreatedAt = result[0].CreatedAt
	} else {
		return nil, errors.New(sql.ErrNoRows.Error())
	}

	for i, book := range result {
		tempBook := dto.GetBookDTO{
			ID:          book.ID_2,
			BookName:    book.BookName,
			BookDesc:    book.BookDesc,
			AuthorID:    book.AuthorID,
			PublishDate: book.PublishDate,
			CreatedAt:   book.CreatedAt,
		}
		books[i] = tempBook

		tempManagment := &dto.GetBookManagmentDTO{
			ID:            book.ID_3.Int64,
			BookID:        book.ID_2,
			TotalQuantity: book.TotalQuantity.Int32,
			TotalInLab:    book.TotalInLab.Int32,
			TotalOutLab:   book.TotalOutLab.Int32,
			CreatedAt:     book.CreatedAt_3.Time,
		}

		if tempManagment.BookID == tempBook.ID {
			managemet = append(managemet, *tempManagment)
			books[i].BookManagment = managemet
		}

	}

	author.Books = books

	return author, nil
}
