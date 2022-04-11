package db

import (
	"context"
	"go-lang-app/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateQuiz(t *testing.T) {
	quiz := createRandomQuiz(t)
	require.NotNil(t, quiz)
	require.NotZero(t, quiz.ID)
}

func TestUpdateQuiz(t *testing.T) {
	quiz := createRandomQuiz(t)
	require.NotNil(t, quiz)
	require.NotZero(t, quiz.ID)

	ctx := context.Background()
	name := util.RandomString(6)
	description := util.RandomString(6)

	// Update a quiz
	arg := &UpdateQuizParams{
		ID:          quiz.ID,
		Name:        name,
		Description: description,
	}

	err := testQuery.UpdateQuiz(ctx, *arg)
	require.NoError(t, err)
	require.NotNil(t, quiz)
	require.NotZero(t, quiz.ID)
	require.Equal(t, arg.ID, quiz.ID)
}


func TestDeleteQuiz(t *testing.T) {
	quiz := createRandomQuiz(t)
	require.NotNil(t, quiz)
	require.NotZero(t, quiz.ID)

	ctx := context.Background()

	// Delete a quiz
	err := testQuery.DeleteQuiz(ctx, quiz.ID)
	require.NoError(t, err)
	require.NotNil(t, quiz)
	require.NotZero(t, quiz.ID)
}

func TestGetAllQuizzes(t *testing.T) {
	ctx := context.Background()
	for i := 0; i < 10; i++ {
		createRandomQuiz(t)
	}

	// Get all quizzes
	args := &GetAllQuizzesParams{
		Limit: 10,
		Offset: 0,

	}

	quizzes, err := testQuery.GetAllQuizzes(ctx, *args)
	require.NoError(t, err)
	require.NotNil(t, quizzes)
	require.NotZero(t, len(quizzes))
}