-- +goose up
CREATE TABLE feeds_follows(
	id UUID PRIMARY KEY,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	user_id UUID NOT NULL references users(id) on delete cascade,
	feed_id UUID NOT NULL references feeds(id) on delete cascade,
	UNIQUE (user_id, feed_id)
);

-- +goose down
drop table feeds_follows;

