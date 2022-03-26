-- name: CreateUser :one
INSERT INTO "user" (
    "first_name",
    "last_name",
    "email",
    "password"
    ) VALUES (
        $1,
        $2,
        $3,
        $4
    ) RETURNING *;

-- name: UpdateUserProfile :exec
UPDATE "user" SET
    "profile_image" = $1,
    "updated_at" = now()
WHERE "id" = $2;

-- name: DeleteUser :exec
UPDATE "user" SET
    "deleted_at" = now()
WHERE "id" = $1;

-- name: FindUser :one
SELECT * FROM "user" WHERE "id" = $1 and "deleted_at" is null LIMIT 1;

-- name: FindUserByEmail :one
SELECT * FROM "user" WHERE "email" = $1 and "deleted_at" is null;

-- name: ChangePassword :exec
UPDATE "user" SET
    "password" = $1,
    "updated_at" = now()
WHERE "id" = $2;


-- name: GetUsers :many
SELECT * FROM "user" WHERE "deleted_at" is null ORDER BY "id" ASC LIMIT $1 OFFSET $2;



