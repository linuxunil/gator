
-- +goose up
CREATE TABLE feeds(
	id UUID PRIMARY KEY,
	created_at timestamp not null,
	updated_at timestamp not null,
	name TEXT NOT NULL,
	url TEXT UNIQUE NOT NULL

);
-- +goose down
DROP TABLE feeds;
