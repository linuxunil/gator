-- +goose up
create table posts (
	id uuid primary key,
	created_at timestamp not null,
	updated_at timestamp,
	title text not null,
	url text not null,
	description text not null,
	published_at timestamp,
	feed_id UUID NOT NULL references feeds(id) on delete cascade
);
-- +goose down
drop table posts;
