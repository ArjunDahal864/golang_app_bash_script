// Code generated by sqlc. DO NOT EDIT.
// source: answer.sql

package db

import (
	"context"
)

const createAnswer = `-- name: CreateAnswer :one
insert into "answer" (
    "question_id",
    "answer",
    "created_by",
    "created_at",
    "updated_at"
    ) values (
        $1,
        $2,
        $3,
        now(),
        now()
    ) RETURNING id, question_id, answer, created_by, created_at, updated_at, deleted_at
`

type CreateAnswerParams struct {
	QuestionID int32  `json:"question_id"`
	Answer     string `json:"answer"`
	CreatedBy  int32  `json:"created_by"`
}

func (q *Queries) CreateAnswer(ctx context.Context, arg CreateAnswerParams) (Answer, error) {
	row := q.db.QueryRowContext(ctx, createAnswer, arg.QuestionID, arg.Answer, arg.CreatedBy)
	var i Answer
	err := row.Scan(
		&i.ID,
		&i.QuestionID,
		&i.Answer,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteAnswer = `-- name: DeleteAnswer :exec
update "answer" set
    "deleted_at" = now()
where "id" = $1
`

func (q *Queries) DeleteAnswer(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteAnswer, id)
	return err
}

const getAllAnswers = `-- name: GetAllAnswers :many
select id, question_id, answer, created_by, created_at, updated_at, deleted_at from "answer" 
where "deleted_at"
is null
order by "id"
asc limit $1 offset $2
`

type GetAllAnswersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetAllAnswers(ctx context.Context, arg GetAllAnswersParams) ([]Answer, error) {
	rows, err := q.db.QueryContext(ctx, getAllAnswers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Answer{}
	for rows.Next() {
		var i Answer
		if err := rows.Scan(
			&i.ID,
			&i.QuestionID,
			&i.Answer,
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

const getAnswerByUserId = `-- name: GetAnswerByUserId :many
select id, question_id, answer, created_by, created_at, updated_at, deleted_at from "answer" 
where "deleted_at"
is null
and "created_by" = $1
order by "id"
asc limit $2 offset $3
`

type GetAnswerByUserIdParams struct {
	CreatedBy int32 `json:"created_by"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *Queries) GetAnswerByUserId(ctx context.Context, arg GetAnswerByUserIdParams) ([]Answer, error) {
	rows, err := q.db.QueryContext(ctx, getAnswerByUserId, arg.CreatedBy, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Answer{}
	for rows.Next() {
		var i Answer
		if err := rows.Scan(
			&i.ID,
			&i.QuestionID,
			&i.Answer,
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

const getAnswersByQuestionId = `-- name: GetAnswersByQuestionId :many
select id, question_id, answer, created_by, created_at, updated_at, deleted_at from "answer" 
where "deleted_at"
is null
and "question_id" = $1
order by "id"
asc limit $2 offset $3
`

type GetAnswersByQuestionIdParams struct {
	QuestionID int32 `json:"question_id"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

func (q *Queries) GetAnswersByQuestionId(ctx context.Context, arg GetAnswersByQuestionIdParams) ([]Answer, error) {
	rows, err := q.db.QueryContext(ctx, getAnswersByQuestionId, arg.QuestionID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Answer{}
	for rows.Next() {
		var i Answer
		if err := rows.Scan(
			&i.ID,
			&i.QuestionID,
			&i.Answer,
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

const updateAnswer = `-- name: UpdateAnswer :exec
update "answer" set
    "answer" = $1,
    "updated_at" = now()
where "id" = $2
`

type UpdateAnswerParams struct {
	Answer string `json:"answer"`
	ID     int32  `json:"id"`
}

func (q *Queries) UpdateAnswer(ctx context.Context, arg UpdateAnswerParams) error {
	_, err := q.db.ExecContext(ctx, updateAnswer, arg.Answer, arg.ID)
	return err
}
