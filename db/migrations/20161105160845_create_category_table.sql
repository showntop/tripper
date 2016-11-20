
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE categories (
  	id serial PRIMARY KEY,
    name varchar,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE categories;

