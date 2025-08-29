-- name: CreateBook :one
INSERT INTO books (id, title, isbn, publication_date, number_of_pages)
VALUES (
    ?, ?, ?, ?, ?
)
RETURNING *;

-- name: ListBooks :many
SELECT * FROM books;

-- name: GetBook :one
SELECT * FROM books WHERE title = ?;

-- name: DeleteBooks :exec
DELETE FROM books;