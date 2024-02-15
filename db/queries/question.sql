-- name: CreateQuestion :one
INSERT INTO questions (
    quiz_id,
    title
) VALUES (
    $1, $2
)
RETURNING *;

-- name: GetQuestion :one
SELECT * FROM questions
WHERE id = $1 LIMIT 1;

-- name: ListQuestions :many
SELECT * from questions
WHERE quiz_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: ListAllQuestions :many
SELECT * from questions
Where quiz_id = $1
ORDER BY id;

-- name: UpdateQuestion :one
UPDATE questions
SET title = $2
WHERE id = $1
RETURNING *;

-- name: DeleteQuestion :exec
DELETE FROM questions
WHERE id = $1;
