// +build e2e

package tests

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreateMovie(t *testing.T) {
	t.Run("should post movie", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().SetHeader("Authorization", "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.2yPNTTY3Y5jUwYBPAJAfUc84Ybv2qPbZY_OHI7tzuug").
		SetBody(`
		{"title": "title", 
		"slug": "slug", 
		"author": "tester", 
		"producer": "producer", 
		"description": "description about movie", 
		"duration": "1s"}`).
		Post("http://localhost:8080/api/v1/movie")

		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode())
	})

	t.Run("shouldn't post movie without JWT", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().SetBody(`
		{"title": "title", 
		"slug": "slug", 
		"author": "tester", 
		"producer": "producer", 
		"description": "description about movie", 
		"duration": "1s"}`).
		Post("http://localhost:8080/api/v1/movie")

		assert.NoError(t, err)
		assert.Equal(t, 401, resp.StatusCode())
	})
}

