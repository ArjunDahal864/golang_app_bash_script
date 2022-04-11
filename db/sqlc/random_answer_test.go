package db

import (
	"context"
	"go-lang-app/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomAnswer(t *testing.T) Answer {
	ctx := context.Background()
	question := createRandomQuestion(t)
	require.NotNil(t, question)
	require.NotZero(t, question.ID)

	user := createRandomUser(t)
	require.NotNil(t, user)
	require.NotZero(t, user.ID)

	answer := util.RandomString(6)

	args := &CreateAnswerParams{
		QuestionID: question.ID,
		Answer:     answer,
		CreatedBy:  user.ID,
	}
	ans, err := testQuery.CreateAnswer(ctx, *args)
	require.NoError(t, err)
	require.NotNil(t, answer)
	require.NotZero(t, ans.ID)
	require.Equal(t, question.ID, ans.QuestionID)
	require.Equal(t, answer, ans.Answer)
	require.Equal(t, user.ID, ans.CreatedBy)
	return ans
}
