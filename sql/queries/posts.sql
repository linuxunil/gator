-- name: CreatePost :one
insert into posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
values ($1,$2,$3,$4,$5,$6,$7,$8)
	returning *;

-- name: GetPostsForUser :many
select * 
from posts 
where posts.feed_id in (
	select feeds.id as feed_id 
	from feeds_follows
	inner join feeds on feeds_follows.feed_id = feeds.id
	inner join users on feeds_follows.user_id = users.id
	where feeds_follows.user_id = $1
) limit $2;
