
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO competitions(name) VALUES ('World Cup 2014');

INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (1, 2, 1, 1, null, null, '2014-06-11 20:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (3, 4, 11, 1, null, null, '2014-06-13 16:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (5, 6, 7, 1, null, null, '2014-06-12 19:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (7, 8, 8, 1, null, null, '2014-06-12 22:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (9, 10, 2, 1, null, null, '2014-06-14 16:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (13, 14, 3, 1, null, null, '2014-06-13 19:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (15, 16, 10, 1, null, null, '2014-06-13 22:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (11, 12, 6, 1, null, null, '2014-06-14 01:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (17, 18, 5, 1, null, null, '2014-06-15 16:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (19, 20, 12, 1, null, null, '2014-06-14 19:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (21, 22, 4, 1, null, null, '2014-06-14 22:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (25, 26, 7, 1, null, null, '2014-06-16 16:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (23, 24, 9, 1, null, null, '2014-06-15 19:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (27, 28, 11, 1, null, null, '2014-06-15 22:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (29, 30, 2, 1, null, null, '2014-06-17 16:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (1, 3, 3, 1, null, null, '2014-06-16 19:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (1, 2, 1, 1, null, null, '2014-06-11 20:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (3, 4, 11, 1, null, null, '2014-06-13 16:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (5, 6, 7, 1, null, null, '2014-06-12 19:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (7, 8, 8, 1, null, null, '2014-06-12 22:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (9, 10, 2, 1, null, null, '2014-06-14 16:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (13, 14, 3, 1, null, null, '2014-06-13 19:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (15, 16, 10, 1, null, null, '2014-06-13 22:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (11, 12, 6, 1, null, null, '2014-06-14 01:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (17, 18, 5, 1, null, null, '2014-06-15 16:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (19, 20, 12, 1, null, null, '2014-06-14 19:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (21, 22, 4, 1, null, null, '2014-06-14 22:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (25, 26, 7, 1, null, null, '2014-06-16 16:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (23, 24, 9, 1, null, null, '2014-06-15 19:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (27, 28, 11, 1, null, null, '2014-06-15 22:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (29, 30, 2, 1, null, null, '2014-06-17 16:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (1, 3, 3, 1, null, null, '2014-06-16 19:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (31, 32, 8, 1, null, null, '2014-06-16 22:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (8, 6, 12, 1, null, null, '2014-06-18 16:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (5, 7, 4, 1, null, null, '2014-06-17 19:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (4, 2, 10, 1, null, null, '2014-06-17 22:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (9, 11, 5, 1, null, null, '2014-06-19 16:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (13, 15, 1, 1, null, null, '2014-06-18 19:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (12, 10, 11, 1, null, null, '2014-06-18 22:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (16, 14, 6, 1, null, null, '2014-06-20 16:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (17, 19, 7, 1, null, null, '2014-06-19 19:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (20, 18, 9, 1, null, null, '2014-06-19 22:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (21, 23, 2, 1, null, null, '2014-06-21 16:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (25, 27, 3, 1, null, null, '2014-06-20 19:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (24, 22, 8, 1, null, null, '2014-06-20 22:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (29, 31, 4, 1, null, null, '2014-06-22 16:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (32, 30, 12, 1, null, null, '2014-06-21 19:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (28, 26, 10, 1, null, null, '2014-06-21 22:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (6, 7, 1, 1, null, null, '2014-06-23 16:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (8, 5, 9, 1, null, null, '2014-06-23 16:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (4, 1, 5, 1, null, null, '2014-06-22 20:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (2, 3, 6, 1, null, null, '2014-06-22 20:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (16, 13, 11, 1, null, null, '2014-06-24 16:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (14, 15, 2, 1, null, null, '2014-06-24 16:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (12, 9, 8, 1, null, null, '2014-06-23 20:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (10, 11, 3, 1, null, null, '2014-06-23 20:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (24, 21, 12, 1, null, null, '2014-06-25 16:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (22, 23, 7, 1, null, null, '2014-06-25 16:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (20, 17, 10, 1, null, null, '2014-06-24 20:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (18, 19, 4, 1, null, null, '2014-06-24 20:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (26, 27, 5, 1, null, null, '2014-06-26 16:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (28, 25, 6, 1, null, null, '2014-06-26 16:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (32, 29, 1, 1, null, null, '2014-06-25 20:00:00 +5:30', 1);
INSERT INTO matches(home_team_id, away_team_id, stadium_id, type, home_team_score, away_team_score, match_time, competition_id) VALUES (30, 31, 9, 1, null, null, '2014-06-25 20:00:00 +5:30', 1);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DELETE FROM competitions;

DELETE FROM matches;
