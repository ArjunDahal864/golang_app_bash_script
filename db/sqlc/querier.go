// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	BlacklistSession(ctx context.Context, id uuid.UUID) error
	ChangePassword(ctx context.Context, arg ChangePasswordParams) error
	CreateAnswer(ctx context.Context, arg CreateAnswerParams) (Answer, error)
	CreateOption(ctx context.Context, arg CreateOptionParams) (Option, error)
	CreateQuestion(ctx context.Context, arg CreateQuestionParams) (Question, error)
	CreateQuiz(ctx context.Context, arg CreateQuizParams) (Quiz, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAnswer(ctx context.Context, id int32) error
	DeleteOption(ctx context.Context, id int32) error
	DeleteQuestion(ctx context.Context, id int32) error
	DeleteQuiz(ctx context.Context, id int32) error
	DeleteUser(ctx context.Context, id int32) error
	FindUser(ctx context.Context, id int32) (User, error)
	FindUserByEmail(ctx context.Context, email string) (User, error)
	GetAllAnswers(ctx context.Context, arg GetAllAnswersParams) ([]Answer, error)
	GetAllOptions(ctx context.Context, arg GetAllOptionsParams) ([]Option, error)
	GetAllQuestions(ctx context.Context, arg GetAllQuestionsParams) ([]Question, error)
	GetAllQuizzes(ctx context.Context, arg GetAllQuizzesParams) ([]Quiz, error)
	GetAnswerByUserId(ctx context.Context, arg GetAnswerByUserIdParams) ([]Answer, error)
	GetAnswersByQuestionId(ctx context.Context, arg GetAnswersByQuestionIdParams) ([]Answer, error)
	GetOptionsByQuestionId(ctx context.Context, arg GetOptionsByQuestionIdParams) ([]Option, error)
	GetQuestionsByQuizId(ctx context.Context, arg GetQuestionsByQuizIdParams) ([]Question, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetUsers(ctx context.Context, arg GetUsersParams) ([]User, error)
	UpdateAnswer(ctx context.Context, arg UpdateAnswerParams) error
	UpdateOption(ctx context.Context, arg UpdateOptionParams) error
	UpdateQuestion(ctx context.Context, arg UpdateQuestionParams) error
	UpdateQuiz(ctx context.Context, arg UpdateQuizParams) error
	UpdateUserProfile(ctx context.Context, arg UpdateUserProfileParams) error
}

var _ Querier = (*Queries)(nil)
