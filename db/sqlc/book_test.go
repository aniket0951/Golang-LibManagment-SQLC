package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/aniket0951/lib_managment/sqlc/utils"
	"github.com/stretchr/testify/require"
)

func createRandomBook(t *testing.T) Book {
	args := CreateBookParams{
		BookName:    utils.RandomBook(),
		BookDesc:    utils.RandomAddress(),
		AuthorID:    1,
		PublishDate: time.Now(),
	}

	book, err := testQueries.CreateBook(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, book)

	require.Equal(t, book.BookName, args.BookName)
	require.Equal(t, book.BookDesc, args.BookDesc)
	require.Equal(t, book.AuthorID, args.AuthorID)

	return book
}

func TestCreateBook(t *testing.T) {
	createRandomBook(t)
}

func TestGetBook(t *testing.T) {

	book, err := testQueries.GetBook(context.Background(), 1)

	require.NoError(t, err)
	require.NotEmpty(t, book)
}

func TestGetBookWithAuthor(t *testing.T) {

	book, err := testQueries.BookWithAuthor(context.Background(), 1)

	require.NoError(t, err)
	require.NotEmpty(t, book)

	fmt.Println("Book And Author : \n", book.AuthorName)
}
