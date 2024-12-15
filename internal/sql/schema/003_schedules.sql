-- +goose Up
CREATE TABLE dashboard (
    date TIMESTAMP NOT NULL,
);

-- +goose Down
DELETE TABLE dashboard;