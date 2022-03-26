package db

import (
	"context"
	"database/sql"
	"go-lang-app/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	user := createRandomUser(t)
	require.NotNil(t, user)
	require.NotZero(t, user.ID)
}

func TestFindUserByID(t *testing.T) {
	ctx := context.Background()
	user := createRandomUser(t)
	user, err := testQuery.FindUser(ctx, user.ID)
	require.NoError(t, err)
	require.NotNil(t, user)
	require.NotZero(t, user.ID)
}

func TestFindUserByEmail(t *testing.T) {
	ctx := context.Background()
	user := createRandomUser(t)
	user, err := testQuery.FindUserByEmail(ctx, user.Email)
	require.NoError(t, err)
	require.NotNil(t, user)
	require.NotZero(t, user.ID)
}

func TestGetUsers(t *testing.T) {
	ctx := context.Background()
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}
	users, err := testQuery.GetUsers(ctx, GetUsersParams{
		Limit:  10,
		Offset: 0,
	})
	require.NoError(t, err)
	require.NotNil(t, users)
	require.NotZero(t, len(users))
}

func TestChangePassword(t *testing.T) {
	ctx := context.Background()
	user := createRandomUser(t)
	hash, err := util.HashPassword(util.RandomString(10))
	require.NoError(t, err)
	require.NotEmpty(t, hash)
	err = testQuery.ChangePassword(ctx, ChangePasswordParams{
		ID:       user.ID,
		Password: hash,
	})
	require.NoError(t, err)
	user, err = testQuery.FindUser(ctx, user.ID)
	require.NoError(t, err)
	require.NotNil(t, user)

}

func TestUpdateProfile(t *testing.T) {
	ctx := context.Background()
	user := createRandomUser(t)
	err := testQuery.UpdateUserProfile(ctx, UpdateUserProfileParams{
		ID: user.ID,
		ProfileImage: sql.NullString{
			String: util.RandomEmail(),
			Valid:  true,
		},
	})
	require.NoError(t, err)
	user, err = testQuery.FindUser(ctx, user.ID)
	require.NoError(t, err)
	require.NotNil(t, user)
}

func TestDeleteUser(t *testing.T){
	ctx := context.Background()
	user := createRandomUser(t)
	err := testQuery.DeleteUser(ctx, user.ID)
	require.NoError(t, err)
	user, err = testQuery.FindUser(ctx, user.ID)
	require.Error(t, err)
	require.Equal(t, "", user.Email)
}