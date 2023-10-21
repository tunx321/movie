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
		Title:       sql.NullString{String: mv.Author, Valid: true},
		Slug:        sql.NullString{String: mv.Slug, Valid: true},
		Description: sql.NullString{String: mv.Description, Valid: true},
		Producer:    sql.NullString{String: mv.Producer, Valid: true},
		Duration:    sql.NullString{String: mv.Duration, Valid: true},
		Author:      sql.NullString{String: mv.Author, Valid: true},
	}

	_, err := d.Client.ExecContext(ctx,
		`INSERT INTO movies (id, title, slug, descript, producer, duration, author) VALUES (:id, :title, :slug, :description, :producer, :duration, :author)`,
		createRow)
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
	cmtRow := MovieRow{
			ID:          id,
			Title:       sql.NullString{String: mv.Author, Valid: true},
			Slug:        sql.NullString{String: mv.Slug, Valid: true},
			Description: sql.NullString{String: mv.Description, Valid: true},
			Producer:    sql.NullString{String: mv.Producer, Valid: true},
			Duration:    sql.NullString{String: mv.Duration, Valid: true},
			Author:      sql.NullString{String: mv.Author, Valid: true},
		
	}
	_, err := d.Client.ExecContext(
		ctx, 
		`UPDATE comments SET
		title = :title,
		slug = :slug,
		descript = :description,
		producer = :producer,
		duration = :duration,
		author = :author
		WHERE id = :id `,
		cmtRow,
	)

	if err != nil{
		return movies.Movie{}, fmt.Errorf("failed to update comment: %w", err) 
	}



	return convertMovieRowToMovie(cmtRow), nil
}