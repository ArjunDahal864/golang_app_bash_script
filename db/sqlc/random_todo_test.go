package db

import (
	"context"
	"database/sql"
	"go-lang-app/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func generateRandomTodo(t *testing.T) Todo {
	usr := createRandomUser(t)
	ctx := context.Background()

	arg := &CreateTodoParams{
		Title:       util.RandomString(6),
		Description: sql.NullString{String: util.RandomString(6), Valid: true},
		IsCompleted: false,
		UserID:      usr.ID,
	}
	todo, err := testQuery.CreateTodo(ctx, *arg)
	require.NoError(t, err)
	require.NotNil(t, todo)
	require.NotZero(t, todo.ID)
	return todo

}
