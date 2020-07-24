-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS super_heroes (
		id bigserial CONSTRAINT pk_super_hero_id primary key,
		uuid uuid NOT NULL,
		super_hero_api_id INT,
		"name" VARCHAR(255) NOT NULL,
		created_at TIMESTAMP,
		updated_at TIMESTAMP
	);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS super_heroes;