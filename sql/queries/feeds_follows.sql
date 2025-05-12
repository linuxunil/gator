
-- name: CreateFeedFollow :one
WITH inserted_feed_follow as (
INSERT INTO feeds_follows (id, created_at, updated_at, user_id, feed_id )
	VALUES ($1,$2,$3,$4,$5)
	RETURNING *
)
select 
	inserted_feed_follow.*, 
	feeds.name as feed_name, 
	users.name as user_name
from inserted_feed_follow
inner join feeds on inserted_feed_follow.feed_id = feeds.id
inner join users on inserted_feed_follow.user_id = users.id;
--
-- name: GetFeedFollowsForUser :many
select feeds_follows.*, feeds.name as feed_name, users.name as user_name
from feeds_follows
inner join feeds on feeds_follows.feed_id = feeds.id
inner join users on feeds_follows.user_id = users.id
where feeds_follows.user_id = $1;
--
-- name: DeleteFeedFollow :exec
delete from feeds_follows 
where user_id = $1 
AND feed_id = $2;
--

