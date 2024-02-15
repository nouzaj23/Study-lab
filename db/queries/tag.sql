-- name: CreateOrGetTag :one
INSERT INTO tags (name)
VALUES ($1)
ON CONFLICT (name) DO UPDATE
    SET name = EXCLUDED.name
RETURNING *;

-- name: AddTagToQuiz :exec
INSERT INTO tags_quizzes (tag_id, quiz_id) VALUES ($1, $2);

-- name: GetTagsForQuiz :many
SELECT t.* FROM tags t
INNER JOIN tags_quizzes tq ON t.id = tq.tag_id
WHERE tq.quiz_id = $1;

-- name: RemoveTagFromQuiz :exec
DELETE FROM tags_quizzes WHERE tag_id = $1 AND quiz_id = $2;

-- name: DeleteTag :exec
DELETE FROM tags WHERE id = $1;

-- name: ListTags :many
SELECT * from tags
ORDER BY id
LIMIT $1
OFFSET $2;