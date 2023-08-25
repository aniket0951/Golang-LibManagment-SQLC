package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/aniket0951/lib_managment/sqlc/db/api/dto"
	db "github.com/aniket0951/lib_managment/sqlc/db/sqlc"
)

type UserService interface {
	AddUser(dto.CreateUserParamsDTO) (dto.GetUsersDTO, error)
	GetUser(int64) (dto.GetUsersDTO, error)
	GetUsers(dto.ListUsersDTO) ([]dto.GetUsersDTO, error)
	UpdateUser(dto.UpdateUserDTO) (dto.GetUsersDTO, error)
	DeleteUser(int64) error
	PurchaseBookForUser(dto.PurchaseBookDTO) error
	ReturnUserBook(req dto.PurchaseBookDTO) error
	GetUserWithBooks(userId int64) (interface{}, error)
}

type userService struct {
	store *db.Store
}

func NewUserService(db *db.Store) UserService {
	return &userService{
		store: db,
	}
}

func (user *userService) AddUser(newUser dto.CreateUserParamsDTO) (dto.GetUsersDTO, error) {
	args := db.CreateUserParams{
		UserName:     newUser.UserName,
		UserEmail:    newUser.UserEmail,
		UserPassword: newUser.UserPassword,
	}

	new_user, err := user.store.CreateUser(context.Background(), args)
	if err != nil {
		return dto.GetUsersDTO{}, err
	}
	return dto.GetUsersDTO{
		ID:        new_user.ID,
		UserName:  new_user.UserName,
		UserEmail: new_user.UserEmail,
		CreatedAt: new_user.CreatedAt,
	}, err
}

func (userSer *userService) GetUser(user_id int64) (dto.GetUsersDTO, error) {
	new_user, err := userSer.store.GetUser(context.Background(), user_id)

	if err != nil {
		return dto.GetUsersDTO{}, err
	}

	return dto.GetUsersDTO{
		ID:        new_user.ID,
		UserName:  new_user.UserName,
		UserEmail: new_user.UserEmail,
		CreatedAt: new_user.CreatedAt,
	}, nil
}

func (userSer *userService) GetUsers(req dto.ListUsersDTO) ([]dto.GetUsersDTO, error) {
	args := db.GetUsersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	result, err := userSer.store.GetUsers(context.Background(), args)

	if err != nil {
		return nil, err
	}

	users := make([]dto.GetUsersDTO, len(result))

	for i, user := range result {
		users[i] = dto.GetUsersDTO{
			ID:        user.ID,
			UserName:  user.UserName,
			UserEmail: user.UserEmail,
			CreatedAt: user.CreatedAt,
		}
	}

	return users, nil
}

func (userSer *userService) UpdateUser(req dto.UpdateUserDTO) (dto.GetUsersDTO, error) {
	args := db.UpdateUserParams{
		ID:       req.Id,
		UserName: req.UserName,
	}

	user, err := userSer.store.UpdateUser(context.Background(), args)

	if err != nil {
		return dto.GetUsersDTO{}, err
	}

	return dto.GetUsersDTO{
		ID:        user.ID,
		UserName:  user.UserName,
		UserEmail: user.UserEmail,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (userSer *userService) DeleteUser(userId int64) error {

	return userSer.store.DeleteUser(context.Background(), userId)
}

func (userSer *userService) PurchaseBookForUser(req dto.PurchaseBookDTO) error {

	// check is the book available or not first
	book_managment, err := userSer.store.GetBookManagmentByBookID(context.Background(), req.BookId)

	if err != nil {
		return err
	}

	if book_managment.TotalQuantity == book_managment.TotalOutLab {
		return errors.New("currently book is not available")
	}

	_, err = userSer.store.PurchaseBook(context.Background(), req.BookId)

	if err != nil {
		if strings.Contains(err.Error(), "violates check constraint") {
			return errors.New("someting went wrong")
		}
		return err
	}

	args := db.AddUserBooksParams{
		UserID: req.UserId,
		BookID: req.BookId,
		PurchaseDate: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	_, err = userSer.store.AddUserBooks(context.Background(), args)

	if err != nil {
		return err
	}

	return nil
}

func (userSer *userService) ReturnUserBook(req dto.PurchaseBookDTO) error {
	// update first book_managment
	_, err := userSer.store.UpdateUserBookManagment(context.Background(), req.BookId)

	if err != nil {
		return err
	}

	// update user books return_book_date
	args := db.UpdateUserBooksParams{
		UserID: req.UserId,
		BookID: req.BookId,
		BookReturnDate: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	_, err = userSer.store.UpdateUserBooks(context.Background(), args)
	return err
}

func (userSer *userService) GetUserWithBooks(userId int64) (interface{}, error) {
	result, err := userSer.store.GetUserWithBooks(context.Background(), userId)

	if err != nil {
		return nil, err
	}

	user_books := dto.GetUsersDTO{}

	if len(result) > 0 {
		user_books.ID = result[0].ID
		user_books.UserName = result[0].UserName
		user_books.UserEmail = result[0].UserEmail
		user_books.CreatedAt = result[0].CreatedAt
	}

	books := make([]dto.GetBookDTO, 0, len(result))

	for _, user_book := range result {
		book := dto.GetBookDTO{
			ID:          user_book.ID_3,
			BookName:    user_book.BookName,
			BookDesc:    user_book.BookDesc,
			AuthorID:    user_book.AuthorID,
			PublishDate: user_book.PublishDate,
			Author: &dto.GetAuthor{
				ID:            user_book.ID_4,
				AuthorName:    user_book.AuthorName,
				AuthorAddress: user_book.AuthorAddress,
				CreatedAt:     user_book.CreatedAt_4,
			},
		}

		if user_book.ID_5.Valid {
			book.Author.AuthorDetailAddress = &dto.AuthorAddressDTO{
				ID:             user_book.ID_5.Int64,
				AddressLineOne: user_book.AddressLineOne.String,
				City:           user_book.City.String,
				State:          user_book.State.String,
				Country:        user_book.Country.String,
				CreatedAt:      user_book.CreatedAt_5.Time,
			}
		}

		books = append(books, book)
	}

	user_books.Books = append(user_books.Books, books...)

	fmt.Printf("%+v", result)

	return user_books, nil
}
