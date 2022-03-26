create table "todo" (
  "id" serial primary key,
  "title" varchar(500) not null,
  "description" varchar(1000),
  "is_completed" boolean not null default false,
  "user_id" integer not null,
  "created_at" timestamp not null default now(),
  "updated_at" timestamp not null default now(),
  "deleted_at" timestamp null
);

-- name: CreateTodo :one
INSERT INTO "todo" (
    "title",
    "description",
    "is_completed",
    "user_id"
    ) VALUES (
        $1,
        $2,
        $3,
        $4
    ) RETURNING *;

-- name: UpdateTodo :exec
UPDATE "todo" SET
    "title" = $1,
    "description" = $2,
    "is_completed" = $3,
    "updated_at" = now()
WHERE "id" = $4;

-- name: DeleteTodo :exec
UPDATE "todo" SET
    "deleted_at" = now()
WHERE "id" = $1;

-- name: GetTodos :many
SELECT * FROM "todo" 
WHERE "deleted_at" is null and
"user_id" = $1
ORDER BY "id" ASC
LIMIT $2 OFFSET $3;

-- name: GetTodo :one
SELECT * FROM "todo" 
WHERE "deleted_at" is null and
"id" = $1;

-- name: GetTodosCount :one
SELECT count(*) FROM "todo" 
WHERE "deleted_at" is null and
"user_id" = $1;

