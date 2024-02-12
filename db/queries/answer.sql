-- name: CreateAnswer :one
INSERT INTO answers (
    question_id,
    text,
    is_correct
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: GetAnswer :one
SELECT * FROM answers
WHERE id = $1 LIMIT 1;

-- name: ListAnswers :many
SELECT * FROM answers
WHERE question_id = $1
ORDER BY id;

-- name: UpdateAnswer :one
UPDATE answers
SET (text, is_correct) = ($2, $3)
WHERE id = $1
RETURNING *;

-- name: DeleteAnswer :exec
DELETE from answers
WHERE id = $1;