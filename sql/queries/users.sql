-- name: CreateUser :one
select * from users
where name=$1
;