package db

import (
	"context"
	"go-lang-app/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateOption(t *testing.T) {
	opt := createRandomOption(t)
	require.NotNil(t, opt)
	require.NotZero(t, opt.ID)

}

func TestUpdateOption(t *testing.T) {
	ctx := context.Background()
	option := createRandomOption(t)
	require.NotNil(t, option)
	require.NotZero(t, option.ID)

	user := createRandomUser(t)
	require.NotNil(t, user)
	require.NotZero(t, user.ID)
	args := &UpdateOptionParams{
		ID:     option.ID,
		Option: util.RandomString(6),
	}

	err := testQuery.UpdateOption(ctx, *args)
	require.NoError(t, err)
	require.NotNil(t, option)
}

func TestDeleteOption(t *testing.T) {
	ctx := context.Background()
	option := createRandomOption(t)
	require.NotNil(t, option)
	require.NotZero(t, option.ID)

	err := testQuery.DeleteOption(ctx, option.ID)
	require.NoError(t, err)
}

func TestGetOptionsByQuestionID(t *testing.T) {
	var options []Option
	for i := 0; i < 10; i++ {
		opt := createRandomOption(t)
		options = append(options, opt)
		require.NotNil(t, opt)
	}

	args := &GetOptionsByQuestionIdParams{
		QuestionID: options[0].QuestionID,
	}
	ctx := context.Background()
	opts, err := testQuery.GetOptionsByQuestionId(ctx, *args)
	require.NoError(t, err)
	require.NotNil(t, opts)
}

func TestGetAllOptions(t *testing.T) {
	var options []Option
	for i := 0; i < 10; i++ {
		opt := createRandomOption(t)
		options = append(options, opt)
		require.NotNil(t, opt)
	}

	args := &GetAllOptionsParams{
		Limit: 10,
		Offset: 0,
	}
	ctx := context.Background()
	opts, err := testQuery.GetAllOptions(ctx,*args)
	require.NoError(t, err)
	require.NotNil(t, opts)
}





