-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE super_hero_powerstats ADD COLUMN created_at timestamp;
ALTER TABLE super_hero_powerstats ADD COLUMN updated_at timestamp;

ALTER TABLE super_hero_biography ADD COLUMN created_at timestamp;
ALTER TABLE super_hero_biography ADD COLUMN updated_at timestamp;

ALTER TABLE super_hero_images ADD COLUMN created_at timestamp;
ALTER TABLE super_hero_images ADD COLUMN updated_at timestamp;

ALTER TABLE super_hero_appearance ADD COLUMN created_at timestamp;
ALTER TABLE super_hero_appearance ADD COLUMN updated_at timestamp;

ALTER TABLE super_hero_works ADD COLUMN created_at timestamp;
ALTER TABLE super_hero_works ADD COLUMN updated_at timestamp;

ALTER TABLE super_hero_connections ADD COLUMN created_at timestamp;
ALTER TABLE super_hero_connections ADD COLUMN updated_at timestamp;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
