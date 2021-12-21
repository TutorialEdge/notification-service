package models

import "time"

type Notification struct {
	ID          string    `json:"id"`
	Email       string    `json:"email"`
	TemplateID  string    `json:"template_id"`
	Data        string    `json:"data"`
	PublishDate time.Time `json:"publish_date"`
	Delivered   bool      `json:"delivered"`
}
