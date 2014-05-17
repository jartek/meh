
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE teams (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  nick_name VARCHAR(255) NOT NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE teams;
