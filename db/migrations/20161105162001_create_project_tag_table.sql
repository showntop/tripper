
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE project_tags (
  	id serial PRIMARY KEY,
    project_id serial REFERENCES projects (id),
    tag_id serial REFERENCES tags (id),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE project_tags;

