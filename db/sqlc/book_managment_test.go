package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomBookManagement(t *testing.T) BookManagment {
	args := CreateBookManagmentParams{
		BookID:        3,
		TotalQuantity: 10,
		TotalInLab:    10,
		TotalOutLab:   0,
	}

	bookManagment, err := testQueries.CreateBookManagment(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, bookManagment)

	require.Equal(t, bookManagment.BookID, args.BookID)
	require.Equal(t, bookManagment.TotalQuantity, args.TotalQuantity)
	require.Equal(t, bookManagment.TotalInLab, args.TotalInLab)
	require.Equal(t, bookManagment.TotalOutLab, args.TotalOutLab)

	return bookManagment

}

func TestCreateBookManagment(t *testing.T) {
	createRandomBookManagement(t)
}

func TestGetBoookManagment(t *testing.T) {
	bookManagment := createRandomBookManagement(t)

	bookManagment1, err := testQueries.GetBookManagment(context.Background(), bookManagment.ID)
	require.NoError(t, err)
	require.NotEmpty(t, bookManagment1)

	require.Equal(t, bookManagment1.BookID, bookManagment.BookID)
	require.Equal(t, bookManagment1.TotalQuantity, bookManagment.TotalQuantity)
	require.Equal(t, bookManagment1.TotalInLab, bookManagment.TotalInLab)
	require.Equal(t, bookManagment1.TotalOutLab, bookManagment.TotalOutLab)

	require.WithinDuration(t, bookManagment1.CreatedAt, bookManagment.CreatedAt, time.Second)
}

func TestGetBoookManagments(t *testing.T) {
	args := GetBookManagmentsParams{
		Limit:  10,
		Offset: 0,
	}

	bookManagments, err := testQueries.GetBookManagments(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, bookManagments)
}

func TestUpdateBookManagment(t *testing.T) {
	bookManagment1 := createRandomBookManagement(t)

	args := UpdateBookManagmentParams{
		ID:            bookManagment1.ID,
		TotalQuantity: 10,
		TotalInLab:    9,
		TotalOutLab:   1,
	}

	updateBookManagment, err := testQueries.UpdateBookManagment(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, updateBookManagment)

	require.Equal(t, updateBookManagment.TotalQuantity, bookManagment1.TotalQuantity)
	require.WithinDuration(t, updateBookManagment.CreatedAt, bookManagment1.CreatedAt, time.Second)
}

func TestDeleteBookManagment(t *testing.T) {
	bookManagment := createRandomBookManagement(t)

	_ = testQueries.DeleteBookManagment(context.Background(), bookManagment.ID)
}

func TestPurcaseBookFromManagment(t *testing.T) {
	//const user_id = 3
	const book_id = 2

	// book_managment, err := testQueries.GetBookManagmentByBookID(context.Background(), book_id)

	// require.NoError(t, err)
	// require.NotEmpty(t, book_managment)

	update_book_managment, err := testQueries.PurchaseBook(context.Background(), book_id)
	
	require.NoError(t, err)
	require.NotEmpty(t, update_book_managment)

	// require.NotEqual(t, update_book_managment.TotalInLab, book_managment.TotalInLab)
	// require.NotEqual(t, update_book_managment.TotalOutLab, book_managment.TotalOutLab)

}
