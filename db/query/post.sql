-- name: CreatePost :one
INSERT INTO posts (
  title,
  content,
  category,
  image
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetPostById :one
SELECT * FROM posts
WHERE id = $1 LIMIT 1;

-- name: ListPosts :many
SELECT * FROM posts
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdatePost :one
UPDATE posts
set 
title = coalesce(sqlc.narg('title'), title), 
category = coalesce(sqlc.narg('category'), category), 
content = coalesce(sqlc.narg('content'), content) ,
image = coalesce(sqlc.narg('image'), image), 
updated_at = coalesce(sqlc.narg('updated_at '), updated_at ) 
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;

-- name: LikePost :exec
UPDATE posts
set
"like" = "like" + 1
WHERE id = $1;

-- name: DislikePost :exec
UPDATE posts
set
"dislike" = "dislike" + 1
WHERE id = $1;

-- name: GetLike :one
SELECT "like" FROM posts
WHERE id = $1 LIMIT 1;

-- name: GetDislike :one
SELECT "dislike" FROM posts
WHERE id = $1 LIMIT 1;

-- name: CountPosts :one
SELECT COUNT(*) FROM posts;