package db

import (
	"context"
	"go-lang-app/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomQuiz(t *testing.T) Quiz {
	ctx := context.Background()
	name := util.RandomString(6)
	description := util.RandomString(6)
	createdBy := createRandomUser(t)

	// Create a new quiz
	arg := &CreateQuizParams{
		Name:        name,
		Description: description,
		CreatedBy:   createdBy.ID,
	}
	quiz, err := testQuery.CreateQuiz(ctx, *arg)
	require.NoError(t, err)
	require.NotNil(t, quiz)
	require.NotZero(t, quiz.ID)
	require.Equal(t, name, quiz.Name)
	require.Equal(t, description, quiz.Description)
	return quiz
}
