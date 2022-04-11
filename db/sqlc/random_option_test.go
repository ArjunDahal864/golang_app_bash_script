package db

import (
	"context"
	"go-lang-app/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomOption(t *testing.T) Option {
	ctx := context.Background()
	question := createRandomQuestion(t)
	require.NotNil(t, question)
	require.NotZero(t, question.ID)

	user := createRandomUser(t)
	require.NotNil(t, user)
	require.NotZero(t, user.ID)

	option := util.RandomString(6)

	args := &CreateOptionParams{
		QuestionID: question.ID,
		Option:     option,
		CreatedBy:  user.ID,
		IsCorrect:  util.RandomBool(),
	}
	opt, err := testQuery.CreateOption(ctx, *args)
	require.NoError(t, err)
	require.NotNil(t, option)
	require.NotZero(t, opt.ID)
	require.Equal(t, question.ID, opt.QuestionID)
	require.Equal(t, option, opt.Option)
	require.Equal(t, user.ID, opt.CreatedBy)
	return opt
}
