package db

import (
	"context"
	"go-lang-app/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAnswer(t *testing.T) {
	ans := createRandomAnswer(t)
	require.NotNil(t, ans)
	require.NotZero(t, ans.ID)
}

func TestGetAllAnswers(t *testing.T) {
	var anss []Answer
	for i := 0; i < 10; i++ {
		createRandomAnswer(t)
		anss = append(anss, createRandomAnswer(t))
	}

	ctx := context.Background()
	args := &GetAllAnswersParams{
		Limit:  10,
		Offset: 0,
	}

	got, err := testQuery.GetAllAnswers(ctx, *args)
	require.NoError(t, err)
	require.NotNil(t, got)
}

func TestGetAllAnswersByQuestionId(t *testing.T) {
	ctx := context.Background()
	question := createRandomQuestion(t)
	require.NotNil(t, question)
	require.NotZero(t, question.ID)

	for i := 0; i < 10; i++ {
		createRandomAnswer(t)
	}

	args := &GetAnswersByQuestionIdParams{
		QuestionID: question.ID,
		Limit:      10,
		Offset:     0,
	}

	got, err := testQuery.GetAnswersByQuestionId(ctx, *args)
	require.NoError(t, err)
	require.NotNil(t, got)
}

func TestGetAnswerByUserId(t *testing.T) {
	var answers []Answer
	for i := 0; i < 10; i++ {
		answers = append(answers, createRandomAnswer(t))
	}
	args := &GetAnswerByUserIdParams{
		CreatedBy: answers[0].CreatedBy,
		Limit:     10,
		Offset:    0,
	}
	asnwes,err:= testQuery.GetAnswerByUserId(context.Background(), *args)
	require.NoError(t, err)
	require.NotNil(t, asnwes)

}

func TestUpdateAnswer(t *testing.T) {
	ans := createRandomAnswer(t)
	require.NotNil(t, ans)
	require.NotZero(t, ans.ID)

	ctx := context.Background()
	answerText := util.RandomString(6)

	// Update a answer
	arg := &UpdateAnswerParams{
		ID:     ans.ID,
		Answer: answerText,
	}

	err := testQuery.UpdateAnswer(ctx, *arg)
	require.NoError(t, err)
	require.NotNil(t, ans)
	require.NotZero(t, ans.ID)
}

func TestDeleteAnswer(t *testing.T) {
	ans := createRandomAnswer(t)
	require.NotNil(t, ans)
	require.NotZero(t, ans.ID)

	ctx := context.Background()

	// Delete a answer
	err := testQuery.DeleteAnswer(ctx, ans.ID)
	require.NoError(t, err)
}
