package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/aniket0951/lib_managment/sqlc/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAuthor(t *testing.T) Author {
	authorName := utils.RandomAuthor()
	address := utils.RandomAddress()

	args := CreateAuthorParams{
		AuthorName:    authorName,
		AuthorAddress: address,
	}

	author, err := testQueries.CreateAuthor(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, author)

	require.Equal(t, args.AuthorName, author.AuthorName)
	require.Equal(t, args.AuthorAddress, author.AuthorAddress)

	require.NotZero(t, author.ID)
	require.NotZero(t, author.CreatedAt)

	return author
}

func TestCreateAuthor(t *testing.T) {
	createRandomAuthor(t)
}

func TestGetAuthorById(t *testing.T) {
	author1 := createRandomAuthor(t)

	author2, err := testQueries.GetAuthor(context.Background(), author1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, author2)

	require.Equal(t, author2.AuthorName, author1.AuthorName)
	require.Equal(t, author2.AuthorAddress, author1.AuthorAddress)

	require.WithinDuration(t, author1.CreatedAt, author2.CreatedAt, time.Second)

}

func TestUpdateAuthor(t *testing.T) {
	author1 := createRandomAuthor(t)

	args := UpdateAuthorParams{
		ID:            author1.ID,
		AuthorName:    author1.AuthorName,
		AuthorAddress: "Navnath Nagar, Bolhegavon, A.Nagar",
	}

	updatedAuthor, err := testQueries.UpdateAuthor(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, updatedAuthor)

	require.Equal(t, updatedAuthor.AuthorName, author1.AuthorName)
	require.NotEqual(t, updatedAuthor.AuthorAddress, author1.AuthorAddress)

	require.WithinDuration(t, author1.CreatedAt, updatedAuthor.CreatedAt, time.Second)
}

func TestDeleteAuthor(t *testing.T) {
	_ = testQueries.DeleteAuthor(context.Background(), 11)
}

func TestGetAuthors(t *testing.T) {
	args := GetAuthorsParams{
		Limit:  5,
		Offset: 1,
	}

	_, err := testQueries.GetAuthors(context.Background(), args)

	require.NoError(t, err)
}

func TestGetAuthorWithBookAndManagment(t *testing.T) {
	managment, err := testQueries.GetAuthorWithBooksAndManagment(context.Background(), 1)
	require.NoError(t, err)
	fmt.Println(managment)
}

func TestCreateAuthorAddress(t *testing.T) {
	const author_id = 2

	args := CreateAuthorAddressParams{
		AddressLineOne: "Pimple Nilakh",
		City:           "Pune",
		State:          "Maharashtra",
		Country:        "India",
		AuthorID:       author_id,
	}

	auther_address, err := testQueries.CreateAuthorAddress(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, auther_address)

	require.Equal(t, auther_address.AddressLineOne, args.AddressLineOne)
	require.Equal(t, auther_address.City, args.City)
	require.Equal(t, auther_address.State, args.State)
	require.Equal(t, auther_address.Country, args.Country)
	require.Equal(t, auther_address.AuthorID, args.AuthorID)

}
