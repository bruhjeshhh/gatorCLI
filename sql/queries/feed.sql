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