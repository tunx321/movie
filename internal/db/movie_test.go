//go:build integration


package db

import (
	"context"
	"fmt"
	"testing"
	"github.com/tunx321/movie/internal/movies"
	"github.com/stretchr/testify/assert"
)

func TestMovieDatabase(t *testing.T) {
	t.Run("test create movie", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		mv, err := db.CreateMovie(context.Background(), movies.Movie{
			Title:       "title",
			Slug:        "slug",
			Description: "description",
			Producer:    "producer",
			Duration:    "duration",
			Author:      "author",
		})
		assert.NoError(t, err)

		newMv, err := db.GetMovie(context.Background(), mv.ID)
		assert.NoError(t, err)
		assert.Equal(t, "title", newMv.Title)
		fmt.Println("Tesing the creation of movies")
	})

	t.Run("test delete movie", func(t *testing.T){
		db, err := NewDatabase()
		assert.NoError(t, err)

		mv, err := db.CreateMovie(context.Background(), movies.Movie{
			Title:       "new-title",
			Slug:        "new-slug",
			Description: "new-description",
			Producer:    "new-producer",
			Duration:    "new-duration",
			Author:      "new-author",
		})
		assert.NoError(t, err)

		err = db.DeleteMovie(context.Background(), mv.ID)
		assert.NoError(t, err)

		_, err = db.GetMovie(context.Background(), mv.ID)
		assert.Error(t, err)
	})
}
