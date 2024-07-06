-- +goose Up
CREATE TABLE user_passwords (
  id UUID PRIMARY KEY REFERENCES users,
  password TEXT NOT NULL
);

-- +goose Down 
DROP TABLE password;
