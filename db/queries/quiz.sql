-- name: CreateQuiz :one
INSERT INTO quizzes (
    name
) VALUES (
    $1
)
RETURNING *;

-- name: GetQuiz :one
SELECT * FROM quizzes
WHERE id = $1 LIMIT 1;

-- name: ListQuizzes :many
SELECT * FROM quizzes
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateQuiz :one
UPDATE quizzes
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteQuiz :exec
DELETE from quizzes
WHERE id = $1;