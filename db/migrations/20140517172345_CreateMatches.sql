
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE matches (
  id SERIAL PRIMARY KEY,
  home_team_id integer REFERENCES teams (id) NOT NULL,
  away_team_id integer REFERENCES teams (id) NOT NULL,
  stadium_id integer REFERENCES stadiums (id) NOT NULL,
  CHECK (home_team_id != away_team_id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE matches;
