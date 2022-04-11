-- name: CreateAnswer :one
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
    ) RETURNING *;

-- name: UpdateAnswer :exec
update "answer" set
    "answer" = $1,
    "updated_at" = now()
where "id" = $2;

-- name: DeleteAnswer :exec
update "answer" set
    "deleted_at" = now()
where "id" = $1;

-- name: GetAllAnswers :many
select * from "answer" 
where "deleted_at"
is null
order by "id"
asc limit $1 offset $2;

-- name: GetAnswersByQuestionId :many
select * from "answer" 
where "deleted_at"
is null
and "question_id" = $1
order by "id"
asc limit $2 offset $3;

-- name: GetAnswerByUserId :many
select * from "answer" 
where "deleted_at"
is null
and "created_by" = $1
order by "id"
asc limit $2 offset $3;

