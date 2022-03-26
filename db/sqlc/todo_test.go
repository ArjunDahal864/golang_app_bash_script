package db

import (
	"context"
	"go-lang-app/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateTodo(t *testing.T) {
	todo := generateRandomTodo(t)
	require.NotNil(t, todo)
	require.NotZero(t, todo.ID)
}

func TestGetTodoByID(t *testing.T) {
	todo := generateRandomTodo(t)
	ctx := context.Background()
	todo, err := testQuery.GetTodo(ctx, todo.ID)
	require.NoError(t, err)
	require.NotNil(t, todo)
	require.NotZero(t, todo.ID)
}

func TestGetTodos(t *testing.T) {
	for i := 0; i < 10; i++ {
		generateRandomTodo(t)
	}
	ctx := context.Background()
	todos, err := testQuery.GetTodos(ctx, GetTodosParams{
		Limit:  10,
		Offset: 0,
		UserID: 1,
	})
	require.NoError(t, err)
	require.NotEmpty(t, todos)
}

func TestUpdateTodo(t *testing.T) {
	todo := generateRandomTodo(t)
	ctx := context.Background()
	err := testQuery.UpdateTodo(ctx, UpdateTodoParams{
		ID:          todo.ID,
		Title:       util.RandomString(6),
		IsCompleted: true,
		Description: todo.Description,
	})
	require.NoError(t, err)
	require.NotNil(t, todo)
	require.NotZero(t, todo.ID)
}

func TestDeleteTodo(t *testing.T) {
	todo := generateRandomTodo(t)
	ctx := context.Background()
	err := testQuery.DeleteTodo(ctx, todo.ID)
	require.NoError(t, err)
	require.NotNil(t, todo)
	require.NotZero(t, todo.ID)
}

func TestGetTodosCount(t *testing.T){
	for i := 0; i < 10; i++ {
		generateRandomTodo(t)
	}
	ctx := context.Background()
	count, err := testQuery.GetTodosCount(ctx, 1)
	require.NoError(t, err)
	require.NotZero(t, count)
}

