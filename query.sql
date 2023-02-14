-- name: GetList :one
SELECT * FROM list
WHERE list_id = $1 LIMIT 1;

-- name: CreateList :one
INSERT INTO list (
    list_id,
    list_name
) VALUES (
    gen_random_uuid(),
    $1   
) RETURNING *;

-- name: DeleteList :exec
DELETE FROM list 
WHERE list_id = $1;

-- name: CreateNotification :one
INSERT INTO notifications (
    notification_id,
    notification_name,
    html
) VALUES (
    gen_random_uuid(),
    $1,
    $2
) RETURNING *;

-- name: GetNotification :one
SELECT * FROM notifications
WHERE notification_id = $1 LIMIT 1;

-- name: DeleteNotification :exec
DELETE FROM notifications
WHERE notification_id = $1;


-- name: GetSubscribers :many
SELECT * FROM subscribers
LIMIT $1;

-- name: Unsubscribe :exec
UPDATE subscribers
SET is_subscribed = 'f'
WHERE email = $1;

-- name: GetSubscriber :one
SELECT * FROM subscribers
WHERE subscriber_id = $1;

-- name: CreateSubscriber :one
INSERT INTO subscribers (
    subscriber_id,
    email
) VALUES (
    gen_random_uuid(),
    $1
) RETURNING *;