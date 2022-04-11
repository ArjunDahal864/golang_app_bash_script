package db

import (
	"context"
	"go-lang-app/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomQuestion(t *testing.T) Question{
	quiz := createRandomQuiz(t)
	require.NotNil(t, quiz)
	require.NotZero(t, quiz.ID)

	user := createRandomUser(t)
	require.NotNil(t, user)
	require.NotZero(t, user.ID)

	ctx := context.Background()
	question := util.RandomString(6)

	args := &CreateQuestionParams{
		QuizID:    quiz.ID,
		Question:  question,
		CreatedBy: user.ID,
	}
	que, err := testQuery.CreateQuestion(ctx, *args)
	require.NoError(t, err)
	require.NotNil(t, question)
	require.NotZero(t, que.ID)
	require.Equal(t, quiz.ID, que.QuizID)
	require.Equal(t, question, que.Question)
	require.Equal(t, user.ID, que.CreatedBy)
	return que
}