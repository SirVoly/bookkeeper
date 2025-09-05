-- name: CreateTag :one
INSERT INTO tags (id, name)
VALUES (?, ?)
RETURNING *;

-- name: ListTags :many
SELECT * FROM tags;

-- name: GetTag :one
SELECT * FROM tags WHERE name = ?;

-- name: DeleteTags :exec
DELETE FROM tags;

-- Books to Tags many-to-many relationship queries
-- name: AddTagToBook :one
INSERT INTO book_tags (book_id, tag_id)
VALUES (?, ?)
RETURNING *;

-- name: ListTagsForBook :many
SELECT t.name
FROM tags t
JOIN book_tags bt ON t.id = bt.tag_id
WHERE bt.book_id = ?;

-- name: ListBooksForTag :many
SELECT b.*
FROM books b
JOIN book_tags bt ON b.id = bt.book_id
WHERE bt.tag_id = ?;

-- name: RemoveTagFromBook :exec
DELETE FROM book_tags
WHERE book_id = ? AND tag_id = ?;

-- name: RemoveAllTagsFromBook :exec
DELETE FROM book_tags
WHERE book_id = ?;

-- name: RemoveAllBooksFromTag :exec
DELETE FROM book_tags
WHERE tag_id = ?;

-- name: ClearBookTags :exec
DELETE FROM book_tags;