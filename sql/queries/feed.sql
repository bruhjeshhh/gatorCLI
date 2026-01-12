-- name: Feed_intoDb :one
INSERT INTO feeds(id, created_at, updated_at, name,url,user_id)
 VALUES ($1, $2, $3, $4,$5,$6)
 RETURNING *;

-- name: GetFeed :many
select feeds.name,feeds.url,users.name
from feeds
join users 
on users.id=feeds.user_id
;
-- name: GetFeedby_Id :one 
select name from feeds where id=$1;

-- name: CreateFeedFollow :many
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id,created_at,updated_at,user_id, feed_id)
    VALUES ($1, $2, $3, $4,$5)
    RETURNING *
)
SELECT
    inserted_feed_follow.id,
    inserted_feed_follow.created_at,
    inserted_feed_follow.updated_at,
    inserted_feed_follow.user_id,
    inserted_feed_follow.feed_id,
    users.name AS user_name,
    feeds.name AS feed_name
FROM inserted_feed_follow
JOIN users ON users.id = inserted_feed_follow.user_id
JOIN feeds ON feeds.id = inserted_feed_follow.feed_id;

-- name: GetFeedby_Url :one 
select * from feeds 
where url=$1;


-- name: GetFeedFollowsForUser :many
select feed_id from feed_follows where user_id=$1;

