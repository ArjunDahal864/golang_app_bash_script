// Code generated by sqlc. DO NOT EDIT.
// source: question.sql

package db

import (
	"context"
)

const createQuestion = `-- name: CreateQuestion :one
INSERT INTO "question" (
    "quiz_id",
    "question",
    "created_by",
    "created_at",
    "updated_at"
    ) VALUES (
        $1,
        $2,
        $3,
        now(),
        now()
    ) RETURNING id, quiz_id, question, created_by, created_at, updated_at, deleted_at
`

type CreateQuestionParams struct {
	QuizID    int32  `json:"quiz_id"`
	Question  string `json:"question"`
	CreatedBy int32  `json:"created_by"`
}

func (q *Queries) CreateQuestion(ctx context.Context, arg CreateQuestionParams) (Question, error) {
	row := q.db.QueryRowContext(ctx, createQuestion, arg.QuizID, arg.Question, arg.CreatedBy)
	var i Question
	err := row.Scan(
		&i.ID,
		&i.QuizID,
		&i.Question,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteQuestion = `-- name: DeleteQuestion :exec
UPDATE "question" SET
    "deleted_at" = now()
WHERE "id" = $1
`

func (q *Queries) DeleteQuestion(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteQuestion, id)
	return err
}

const getAllQuestions = `-- name: GetAllQuestions :many
SELECT id, quiz_id, question, created_by, created_at, updated_at, deleted_at FROM "question" 
WHERE "deleted_at"
is null
ORDER BY "id"
ASC LIMIT $1 OFFSET $2
`

type GetAllQuestionsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetAllQuestions(ctx context.Context, arg GetAllQuestionsParams) ([]Question, error) {
	rows, err := q.db.QueryContext(ctx, getAllQuestions, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Question{}
	for rows.Next() {
		var i Question
		if err := rows.Scan(
			&i.ID,
			&i.QuizID,
			&i.Question,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getQuestionsByQuizId = `-- name: GetQuestionsByQuizId :many
SELECT id, quiz_id, question, created_by, created_at, updated_at, deleted_at FROM "question" 
WHERE "deleted_at"
is null
AND "quiz_id" = $1
ORDER BY "id"
ASC LIMIT $2 OFFSET $3
`

type GetQuestionsByQuizIdParams struct {
	QuizID int32 `json:"quiz_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetQuestionsByQuizId(ctx context.Context, arg GetQuestionsByQuizIdParams) ([]Question, error) {
	rows, err := q.db.QueryContext(ctx, getQuestionsByQuizId, arg.QuizID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Question{}
	for rows.Next() {
		var i Question
		if err := rows.Scan(
			&i.ID,
			&i.QuizID,
			&i.Question,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateQuestion = `-- name: UpdateQuestion :exec
UPDATE "question" SET
    "question" = $1,
    "updated_at" = now()
WHERE "id" = $2
`

type UpdateQuestionParams struct {
	Question string `json:"question"`
	ID       int32  `json:"id"`
}

func (q *Queries) UpdateQuestion(ctx context.Context, arg UpdateQuestionParams) error {
	_, err := q.db.ExecContext(ctx, updateQuestion, arg.Question, arg.ID)
	return err
}
