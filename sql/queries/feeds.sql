
-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedByURL :one
SELECT * FROM feeds where $1 = url; 
-- name: MarkFeedFetched :exec
UPDATE feeds
SET updated_at = NOW(), last_fetched_at = NOW()
WHERE id = $1;
-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST;

