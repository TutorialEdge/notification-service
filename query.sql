-- name: GetList :one
SELECT * FROM list
WHERE list_id = $1 LIMIT 1;

-- name: CreateList :one
INSERT INTO list (
    list_name
) VALUES (
    $1   
) RETURNING *;

-- name: DeleteList :exec
DELETE FROM list 
WHERE list_id = $1;

-- name: CreateNotification :one
INSERT INTO notifications (
    notification_name,
    html
) VALUES (
    $1,
    $2
) RETURNING *;

-- name: GetNotification :one
SELECT * FROM notifications
WHERE notification_id = $1 LIMIT 1;

-- name: DeleteNotification :exec
DELETE FROM notifications
WHERE notification_id = $1;


-- name: Unsubscribe :exec
UPDATE subscribers
SET is_subscribed = 'f'
WHERE email = $1;

-- name: GetSusbcriber :one
SELECT * FROM subscribers
WHERE subscriber_id = $1;

-- name: CreateSubscriber :one
INSERT INTO subscribers (
    email
) VALUES (
    $1
) RETURNING *;