-- +goose Up
CREATE TABLE users(
    id UUID PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    first_name TEXT UNIQUE NOT NULL,
    last_name TEXT UNIQUE NOT NULL
);

-- +goose Down 
DROP TABLE users;
