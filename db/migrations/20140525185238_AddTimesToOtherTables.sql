
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE teams
ADD COLUMN created_at TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp,
ADD COLUMN updated_at TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp;

ALTER TABLE stadia
ADD COLUMN created_at TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp,
ADD COLUMN updated_at TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE teams
DROP COLUMN created_at,
DROP COLUMN updated_at;


ALTER TABLE stadia
DROP COLUMN created_at,
DROP COLUMN updated_at;
