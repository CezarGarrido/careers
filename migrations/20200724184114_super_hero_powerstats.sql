-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE super_hero_powerstats (
		id bigserial CONSTRAINT pk_super_hero_powerstat_id primary key,
		uuid uuid NOT NULL,
		super_hero_id INT NOT NULL,
		intelligence VARCHAR(255),
		strength VARCHAR(255),
		speed VARCHAR(255),
		durability VARCHAR(255),
		power VARCHAR(255),
		combat VARCHAR(255),
	 FOREIGN KEY (super_hero_id) REFERENCES super_heroes(id) ON DELETE CASCADE
	);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS super_hero_powerstats;