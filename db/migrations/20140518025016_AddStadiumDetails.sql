
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE stadia
ADD COLUMN capacity int NOT NULL,
ADD COLUMN altitude int NOT NULL,
ADD COLUMN year_built int NOT NULL;

INSERT INTO stadia VALUES (1, 'Arena De Sao Paulo', 'Sao Paulo', 65807, 792, 2014);
INSERT INTO stadia VALUES (2, 'Estadio Mineirao', 'Belo Horizonte', 62547, 800, 1965);
INSERT INTO stadia VALUES (3, 'Estadio Castelao', 'Fortaleza', 64846, 0, 1973);
INSERT INTO stadia VALUES (4, 'Estadio Do Maracana', 'Rio de Janeiro', 76804, 0, 1950);
INSERT INTO stadia VALUES (5, 'Estadio Nacional', 'Brasilia', 68009, 1172, 1974);
INSERT INTO stadia VALUES (6, 'Arena Pernambuco', 'Recife', 44248, 0, 2013);
INSERT INTO stadia VALUES (7, 'Arena Fonte Nova', 'Salvador', 48747, 0, 1951);
INSERT INTO stadia VALUES (8, 'Estadio Pantanal', 'Cuiaba', 42968, 165, 2014);
INSERT INTO stadia VALUES (9, 'Estadio Da Baixada', 'Curitiba', 41456, 920, 1914);
INSERT INTO stadia VALUES (10, 'Estadio Amazonia', 'Manaus', 42374, 72, 1970);
INSERT INTO stadia VALUES (11, 'Estadio Das Dunas', 'Natal', 42086, 45, 2013);
INSERT INTO stadia VALUES (12, 'Estadio Beira-Rio', 'Porto Alegre', 48849, 47, 1969);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE stadia
DROP COLUMN capacity,
DROP COLUMN altitude,
DROP COLUMN year_built;

DELETE FROM stadia;
