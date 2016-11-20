
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE posts (
  	id serial PRIMARY KEY,
  	user_id serial REFERENCES users (id),
    project_id serial REFERENCES projects (id),
    content text,
    like_num int,
    comment_num int,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);



-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE posts;

