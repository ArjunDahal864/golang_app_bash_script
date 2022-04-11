package db

import (
	"context"
	"go-lang-app/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateQuestion(t *testing.T) {
	question := createRandomQuestion(t)
	require.NotNil(t, question)
	require.NotZero(t, question.ID)
}

func TestGetQuestionsByQuizId(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomQuestion(t)
	}

	ctx := context.Background()
	quiz := createRandomQuiz(t)
	require.NotNil(t, quiz)
	require.NotZero(t, quiz.ID)

	args := &GetQuestionsByQuizIdParams{
		QuizID: 5,
		Limit:  10,
		Offset: 0,
	}

	questions, err := testQuery.GetQuestionsByQuizId(ctx, *args)
	require.NoError(t, err)
	require.NotNil(t, questions)
	require.NotZero(t, len(questions))
}

func TestUpdateQuestion(t *testing.T) {
	question := createRandomQuestion(t)
	require.NotNil(t, question)
	require.NotZero(t, question.ID)

	ctx := context.Background()
	questionText := util.RandomString(6)

	// Update a question
	arg := &UpdateQuestionParams{
		ID:       question.ID,
		Question: questionText,
	}

	err := testQuery.UpdateQuestion(ctx, *arg)
	require.NoError(t, err)
	require.NotNil(t, question)
	require.NotZero(t, question.ID)
}

func TestDeleteQuestion(t *testing.T) {
	question := createRandomQuestion(t)
	require.NotNil(t, question)
	require.NotZero(t, question.ID)

	ctx := context.Background()

	// Delete a question
	err := testQuery.DeleteQuestion(ctx, question.ID)
	require.NoError(t, err)
	require.NotNil(t, question)
	require.NotZero(t, question.ID)
}

func TestGetAllQuestions(t *testing.T) {
	ctx := context.Background()
	for i := 0; i < 10; i++ {
		createRandomQuestion(t)
	}

	// Get all questions
	args := &GetAllQuestionsParams{
		Limit: 10,
		Offset: 0,

	}

	questions, err := testQuery.GetAllQuestions(ctx, *args)
	require.NoError(t, err)
	require.NotNil(t, questions)
	require.NotZero(t, len(questions))
}



