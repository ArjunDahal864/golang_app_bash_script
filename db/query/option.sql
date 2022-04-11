-- name: CreateOption :one
insert into "option" (
  "question_id",
  "option",
  "is_correct",
  "created_by",
  "created_at",
  "updated_at"
) values (
    $1,
    $2,
    $3,
    $4,
   now(),
   now()
) RETURNING *;

-- name: UpdateOption :exec
update "option" set
  "option" = $1,
  "is_correct" = $2,
  "updated_at" = now()
where "id" = $3;

-- name: DeleteOption :exec
update "option" set
  "deleted_at" = now()
where "id" = $1;

-- name: GetAllOptions :many
select * from "option" 
where "deleted_at"
is null
order by "id"
asc limit $1 offset $2;

-- name: GetOptionsByQuestionId :many
select * from "option" 
where "deleted_at"
is null
and "question_id" = $1
order by "id"
asc limit $2 offset $3;

