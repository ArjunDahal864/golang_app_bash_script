package db

import (
	"context"
	"go-lang-app/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	ctx := context.Background()
	hash, err := util.HashPassword("password")
	require.NoError(t, err)
	require.NotEmpty(t, hash)

	email := util.RandomEmail()

	firstName := util.RandomString(6)
	lastName := util.RandomString(6)

	// Create a new user
	arg := &CreateUserParams{
		FirstName: firstName,
		LastName:  lastName,
		Password:  hash,
		Email:     email,
	}
	user, err := testQuery.CreateUser(ctx, *arg)
	require.NoError(t, err)
	require.NotNil(t, user)
	require.NotZero(t, user.ID)
	return user
}
