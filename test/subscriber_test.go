//go:build e2e
// +build e2e

package test

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestSubscriptions(t *testing.T) {
	client := resty.New()
	t.Run("Add a new subscriber", func(t *testing.T) {
		resp, err := client.R().
			SetBody(`{"email":"test@test.com"}`).
			Post(BASE_URL + "/api/v1/subscribe")

		assert.Nil(t, err)
		fmt.Println(resp)
	})

	t.Run("Unsubscribe a subscriber", func(t *testing.T) {
		resp, err := client.R().
			SetBody(`{"email":"subscriber@test.com"}`).
			Post(BASE_URL + "/api/v1/subscribe")

		assert.Nil(t, err)
		fmt.Println(resp)

		resp, err = client.R().
			Get(BASE_URL + "/api/v1/unsubscribe?email=subscriber@test.com")

		assert.Nil(t, err)
		fmt.Println(resp)

	})
}