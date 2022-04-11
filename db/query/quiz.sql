-- name: CreateQuiz :one
INSERT INTO "quiz" (
    "name",
    "description",
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

-- name: UpdateQuiz :exec
UPDATE "quiz" SET
    "name" = $1,
    "description" = $2,
    "updated_at" = now()
WHERE "id" = $3;

-- name: DeleteQuiz :exec
UPDATE "quiz" SET
    "deleted_at" = now()
WHERE "id" = $1;

-- name: GetAllQuizzes :many
SELECT * FROM "quiz" 
WHERE "deleted_at" 
is null
ORDER BY "id" 
ASC LIMIT $1 OFFSET $2;
