
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users (
  	id serial PRIMARY KEY,
    username varchar NOT NULL UNIQUE,
    name varchar,
    email varchar UNIQUE,
    mobile varchar UNIQUE,
    hashed_password varchar,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;
