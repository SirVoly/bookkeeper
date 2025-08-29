-- +goose Up
CREATE TABLE books (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    isbn INTEGER NOT NULL,
    publication_date TEXT,
    number_of_pages INTEGER
);


-- +goose Down
DROP TABLE books;