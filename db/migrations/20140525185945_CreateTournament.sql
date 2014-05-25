
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE tournament (
  id SERIAL PRIMARY KEY,
  name varchar NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp
);

ALTER TABLE matches
ADD COLUMN tournament_id int NOT NULL;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE tournament;

ALTER TABLE matches
DROP COLUMN tournament_id;
