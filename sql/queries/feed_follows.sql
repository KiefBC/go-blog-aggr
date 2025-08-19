-- name: CreateFeedFollow :one
WITH new_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT new_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM new_follow
INNER JOIN feeds ON new_follow.feed_id = feeds.id
INNER JOIN users ON new_follow.user_id = users.id;
