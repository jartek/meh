
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE matches
ADD COLUMN match_time TIMESTAMP WITH TIME ZONE,
ADD COLUMN created_at TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp,
ADD COLUMN updated_at TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE matches
DROP COLUMN match_time,
DROP COLUMN created_at,
DROP COLUMN updated_at;
