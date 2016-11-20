
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE projects (
  	id serial PRIMARY KEY,
    name varchar NOT NULL,
    description text,
    version varchar,
    size integer,
    dlink varchar,
    logo_url varchar,
    category_id serial REFERENCES categories (id),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE projects;

