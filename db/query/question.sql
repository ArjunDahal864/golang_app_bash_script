-- name: CreateQuestion :one
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
    ) RETURNING *;

-- name: UpdateQuestion :exec
UPDATE "question" SET
    "question" = $1,
    "updated_at" = now()
WHERE "id" = $2;

-- name: DeleteQuestion :exec
UPDATE "question" SET
    "deleted_at" = now()
WHERE "id" = $1;

-- name: GetAllQuestions :many
SELECT * FROM "question" 
WHERE "deleted_at"
is null
ORDER BY "id"
ASC LIMIT $1 OFFSET $2;

-- name: GetQuestionsByQuizId :many
SELECT * FROM "question" 
WHERE "deleted_at"
is null
AND "quiz_id" = $1
ORDER BY "id"
ASC LIMIT $2 OFFSET $3;