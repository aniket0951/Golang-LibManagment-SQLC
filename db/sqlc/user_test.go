package db

import (
	"context"
	"testing"
	"time"

	"github.com/aniket0951/lib_managment/sqlc/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) *Users {
	args := CreateUserParams{
		UserName: utils.RandomString(5),
		UserEmail: utils.RandomEmail(),
		UserPassword: "123456",
	}

	users, err := testQueries.CreateUser(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, users)

	require.Equal(t, users.UserName, args.UserName)
	require.Equal(t, users.UserEmail, args.UserEmail)
	require.Equal(t, users.UserPassword, args.UserPassword)

	require.NotZero(t, users.ID)

	return &users
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)

	new_user, err := testQueries.GetUser(context.Background(), user.ID)

	require.NoError(t,err)
	require.NotEmpty(t,new_user)

	require.Equal(t, new_user.UserName, user.UserName)
	require.Equal(t, new_user.UserEmail, user.UserEmail)
	require.Equal(t, new_user.UserPassword, user.UserPassword)

	require.WithinDuration(t, user.CreatedAt, new_user.CreatedAt, time.Second)
}

func TestGetUsers(t *testing.T) {
	args := GetUsersParams{
		Limit: 10,
		Offset: 1,
	}

	_, err := testQueries.GetUsers(context.Background(), args)

	require.NoError(t, err)
}

func TestUpdateUser(t *testing.T) {
	user := createRandomUser(t)

	update_args := UpdateUserParams{
		ID: user.ID,
		UserName: "Aniket-Suryawanshi",
	}

	updated_user, err := testQueries.UpdateUser(context.Background(), update_args)

	require.NoError(t,err)
	require.NotEmpty(t, updated_user)

	require.Equal(t, updated_user.ID, user.ID)
	require.NotEqual(t, updated_user.UserName, user.UserName)

	require.WithinDuration(t, updated_user.CreatedAt, user.CreatedAt, time.Second)
}

func TestDeleteUser(t *testing.T) {
	err := testQueries.DeleteUser(context.Background(), 1)
	require.Empty(t, err)
}