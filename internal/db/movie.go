package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/tunx321/movie/internal/movies"
)

type MovieRow struct {
	ID          string
	Title       sql.NullString
	Slug        sql.NullString
	Description sql.NullString
	Producer    sql.NullString
	Duration    sql.NullString
	Author      sql.NullString
}

func convertMovieRowToMovie(m MovieRow) movies.Movie {
	return movies.Movie{
		ID:          m.ID,
		Title:       m.Title.String,
		Slug:        m.Slug.String,
		Description: m.Description.String,
		Producer:    m.Producer.String,
		Duration:    m.Duration.String,
		Author:      m.Author.String,
	}
}

func (d *Database) GetMovie(ctx context.Context, uuid string) (movies.Movie, error) {
	var mvRow MovieRow
	row := d.Client.QueryRowContext(ctx,
		`SELECT id, title, slug, descript, producer, duration, author
	FROM movies WHERE id = $1`, uuid)

	if err := row.Scan(&mvRow.ID, &mvRow.Title, &mvRow.Slug, &mvRow.Description, &mvRow.Producer, &mvRow.Duration, &mvRow.Author); err != nil {
		return movies.Movie{}, fmt.Errorf("error fetching movie by uuid: %w", err)
	}

	return convertMovieRowToMovie(mvRow), nil
}

func (d *Database) CreateMovie(ctx context.Context, mv movies.Movie) (movies.Movie, error) {
	mv.ID = uuid.New().String()
	createRow := MovieRow{
		ID:          mv.ID,
		Title:       sql.NullString{String: mv.Title, Valid: true},
		Slug:        sql.NullString{String: mv.Slug, Valid: true},
		Description: sql.NullString{String: mv.Description, Valid: true},
		Producer:    sql.NullString{String: mv.Producer, Valid: true},
		Duration:    sql.NullString{String: mv.Duration, Valid: true},
		Author:      sql.NullString{String: mv.Author, Valid: true},
	}

	err := d.Client.QueryRowContext(ctx,
		`INSERT INTO movies (id, title, slug, descript, producer, duration, author) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, title, slug, descript, producer, duration, author`,
		createRow.ID, createRow.Title, createRow.Slug, createRow.Description, createRow.Producer, createRow.Duration, createRow.Author).Scan(&mv.ID, &mv.Title, &mv.Slug, &mv.Description, &mv.Producer, &mv.Duration, &mv.Author)
	
	if err != nil {
		return movies.Movie{}, fmt.Errorf("failed to insert movie: %w", err)
	}


	return mv, nil
}


func (d *Database) DeleteMovie(ctx context.Context, id string) error{
	_, err := d.Client.ExecContext(ctx, 
	`DELETE FROM movies WHERE id = $1`, id)

	if err != nil{
		return fmt.Errorf("failed to delete movie: %w", err)
	}

	return nil
}

func (d *Database) UpdateMovie(ctx context.Context, id string, mv movies.Movie) (movies.Movie, error){
	mvRow := MovieRow{
			ID:          id,
			Title:       sql.NullString{String: mv.Author, Valid: true},
			Slug:        sql.NullString{String: mv.Slug, Valid: true},
			Description: sql.NullString{String: mv.Description, Valid: true},
			Producer:    sql.NullString{String: mv.Producer, Valid: true},
			Duration:    sql.NullString{String: mv.Duration, Valid: true},
			Author:      sql.NullString{String: mv.Author, Valid: true},
		
	}
	err := d.Client.QueryRowContext(
		ctx, 
		`UPDATE movies SET
		title = $1,
		slug = $2,
		descript = $3,
		producer = $4,
		duration = $5,
		author = $6
		WHERE id = $7 RETURNING title, slug, descript, producer, duration, author`,
		mvRow.Title, mvRow.Slug, mvRow.Description, mvRow.Producer, mvRow.Duration, mvRow.Author, mvRow.ID,
	).Scan(&mvRow.Title, &mvRow.Slug, &mvRow.Description, &mvRow.Producer, &mvRow.Duration, &mvRow.Author)

	if err != nil{
		return movies.Movie{}, fmt.Errorf("failed to update movies: %w", err) 
	}



	return convertMovieRowToMovie(mvRow), nil
}