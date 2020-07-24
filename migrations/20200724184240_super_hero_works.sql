-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE super_hero_works (
		id bigserial CONSTRAINT pk_super_hero_work_id primary key,
		uuid uuid NOT NULL,
		super_hero_id INT NOT NULL,
		occupation TEXT,
		base TEXT,
		FOREIGN KEY (super_hero_id) REFERENCES super_heroes(id) ON DELETE CASCADE
	);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS super_hero_works;