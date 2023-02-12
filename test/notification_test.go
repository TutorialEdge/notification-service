//go:build e2e
// +build e2e

package test

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

const (
	BASE_URL = "http://localhost:8080"
)

func TestNotifications(t *testing.T) {
	client := resty.New()
	t.Run("create a new notification", func(t *testing.T) {
		resp, err := client.R().
			SetBody(`{"slug": "/", "author": "12345", "body": "hello world"}`).
			Post(BASE_URL + "/api/v1/notification")

		assert.Nil(t, err)
		fmt.Println(resp)

		resp, err = client.R().
			SetBody(`{"slug": "/", "author": "12345", "body": "hello world"}`).
			Get(BASE_URL + "/api/v1/notification/" + "uuid")

		assert.Nil(t, err)
		fmt.Println(resp)
	})

	t.Run("delete a notification", func(t *testing.T) {
		resp, err := client.R().
			SetBody(`{"slug": "/", "author": "12345", "body": "hello world"}`).
			Post(BASE_URL + "/api/v1/notification")

		assert.Nil(t, err)
		fmt.Println(resp)

		resp, err = client.R().
			SetBody(`{"slug": "/", "author": "12345", "body": "hello world"}`).
			Delete(BASE_URL + "/api/v1/notification/" + "todo")

		assert.Nil(t, err)
		fmt.Println(resp)

	})

	t.Run("update a notification", func(t *testing.T) {
		resp, err := client.R().
			SetBody(`{"slug": "/", "author": "12345", "body": "hello world"}`).
			Post(BASE_URL + "/api/v1/notification")

		assert.Nil(t, err)
		fmt.Println(resp)

		resp, err = client.R().
			SetBody(`{"slug": "/", "author": "12345", "body": "hello world"}`).
			Put(BASE_URL + "/api/v1/notification")

		assert.Nil(t, err)
		fmt.Println(resp)

	})
}
