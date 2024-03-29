// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package store

import (
	"database/sql"

	"github.com/google/uuid"
)

type List struct {
	ListID    uuid.UUID
	ListName  string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}

type Notification struct {
	NotificationID   uuid.UUID
	NotificationName sql.NullString
	Html             sql.NullString
	CreatedAt        sql.NullTime
	UpdatedAt        sql.NullTime
	DeletedAt        sql.NullTime
}

type Subscriber struct {
	SubscriberID uuid.UUID
	Email        string
	IsSubscribed sql.NullBool
	CreatedAt    sql.NullTime
	UpdatedAt    sql.NullTime
	DeletedAt    sql.NullTime
}
