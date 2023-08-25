package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPurchaseBookFromLib(t *testing.T) {
	const user_id = 3
	const book_id = 3

	_, err := testQueries.PurchaseBook(context.Background(), book_id)

	require.NoError(t, err)

	args := AddUserBooksParams{
		UserID: user_id,
		BookID: book_id,
		PurchaseDate: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}
	user_book, user_book_err := testQueries.AddUserBooks(context.Background(), args)

	require.NoError(t, user_book_err)
	require.NotEmpty(t, user_book)

	require.Equal(t, user_book.UserID, args.UserID)
	require.Equal(t, user_book.BookID, args.BookID)
	require.NotZero(t, user_book.ID)
}

func TestReturnBookForLib(t *testing.T) {
	const book_id = 3
	const user_id = 3

	book, err := testQueries.UpdateUserBookManagment(context.Background(), book_id)

	require.NoError(t, err)
	require.NotEmpty(t, book)

	// update user_book update return_book_date
	args := UpdateUserBooksParams{
		UserID: user_id,
		BookID: book_id,
		BookReturnDate: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	user_book, err := testQueries.UpdateUserBooks(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, user_book)

	require.Equal(t, user_book.UserID, args.UserID)
	require.Equal(t, user_book.BookID, args.BookID)
}
