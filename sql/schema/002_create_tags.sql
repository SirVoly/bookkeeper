-- +goose Up
CREATE TABLE tags (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE tags;