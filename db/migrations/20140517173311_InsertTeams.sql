
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO teams VALUES(1, 'Brazil', 'Canarinho');
INSERT INTO teams VALUES(2, 'Croatia', 'Vatreni');
INSERT INTO teams VALUES(3, 'Mexico', 'El Tri');
INSERT INTO teams VALUES(4, 'Cameroon', 'Les Lions Indomptables');
INSERT INTO teams VALUES(5, 'Spain', 'La Furia Roja');
INSERT INTO teams VALUES(6, 'Netherlands', 'Clockwork Orange');
INSERT INTO teams VALUES(7, 'Chile', 'La Roja');
INSERT INTO teams VALUES(8, 'Australia', 'Socceroos');
INSERT INTO teams VALUES(9, 'Colombia', 'Los Cafeteros');
INSERT INTO teams VALUES(10, 'Greece', 'Piratiko');
INSERT INTO teams VALUES(11, E'Cote d\'Ivoire', 'Les Elephants');
INSERT INTO teams VALUES(12, 'Japan', 'Blue Samurai');
INSERT INTO teams VALUES(13, 'Uruguay', 'La Celeste');
INSERT INTO teams VALUES(14, 'Costa Rica', 'Los Ticos');
INSERT INTO teams VALUES(15, 'England', 'Three Lions');
INSERT INTO teams VALUES(16, 'Italy', 'Gli Azzuri');
INSERT INTO teams VALUES(17, 'Switzerland', 'The Schweizer Nati');
INSERT INTO teams VALUES(18, 'Ecuador', 'La Tri');
INSERT INTO teams VALUES(19, 'France', 'Les Bleus');
INSERT INTO teams VALUES(20, 'Honduras', 'Los Catrachos');
INSERT INTO teams VALUES(21, 'Argentina', 'La Albicelestes');
INSERT INTO teams VALUES(22, 'Bosnia-Herzegovina', 'Zmajevi');
INSERT INTO teams VALUES(23, 'Iran', 'Team Melli');
INSERT INTO teams VALUES(24, 'Nigeria', 'Super Eagles');
INSERT INTO teams VALUES(25, 'Germany', 'Nationalmannschaft');
INSERT INTO teams VALUES(26, 'Portugal', 'Selecao Das Quinas');
INSERT INTO teams VALUES(27, 'Ghana', 'Black Stars');
INSERT INTO teams VALUES(28, 'United States', 'The Yanks');
INSERT INTO teams VALUES(29, 'Belgium', 'Red Devils');
INSERT INTO teams VALUES(30, 'Algeria', 'Les Fennecs');
INSERT INTO teams VALUES(31, 'Russia', 'Sbornaya');
INSERT INTO teams VALUES(32, 'South Korea', 'Taeguk');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DELETE FROM teams;
