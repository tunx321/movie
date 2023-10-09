package main

import (
	"context"
	"fmt"

	"github.com/tunx321/movie/internal/db"
	"github.com/tunx321/movie/internal/movies"
)

func Run() error {
	fmt.Println("starting up our appplication")
	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("failed to ping to database")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("dailed to migrate database")
		return err
	}
	mvService := movies.NewService(db)

	_, err = mvService.CreateMovie(context.Background(), 
	movies.Movie{
		Title: "testing2",
		Slug: "test2",
		Producer: "Almat2",
		Author: "Alma2t",
		Description: "Testing database2",
	},
	)
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(mvService.GetMovie(context.Background(), "3654bcbf-57b6-4e66-86b5-87fb9790f8a9"))
	fmt.Println("successfully connected and pinged database")
	return nil
}

func main() {
	fmt.Println("Go REST API movie")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
