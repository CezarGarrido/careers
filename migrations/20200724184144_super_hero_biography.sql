-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE super_hero_biography (
		id bigserial CONSTRAINT pk_super_hero_biography_id primary key,
		uuid uuid NOT NULL,
		super_hero_id INT NOT NULL,
		fullname VARCHAR(255),
		alter_egos VARCHAR(255),
		aliases TEXT,
		place_of_birth VARCHAR(255),
		first_appearance VARCHAR(255),
		publisher VARCHAR(255),
		alignment VARCHAR(4) NOT NULL,
		FOREIGN KEY (super_hero_id) REFERENCES super_heroes(id) ON DELETE CASCADE
	);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS super_hero_biography;