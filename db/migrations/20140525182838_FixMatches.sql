
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE matches
ADD COLUMN type int NOT NULL,
ADD COLUMN home_team_score int,
ADD COLUMN away_team_score int;


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE matches
DROP COLUMN type,
DROP COLUMN home_team_score,
DROP COLUMN away_team_score;
