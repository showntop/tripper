
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE post_likes(
	id serial PRIMARY KEY,
	user_id serial REFERENCES users (id),
	post_id serial REFERENCES posts (id),
	created_at timestamp without time zone,
	updated_at timestamp without time zone
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE post_likes;

