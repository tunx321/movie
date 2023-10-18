package main

import (
	"fmt"

	"github.com/tunx321/movie/internal/db"
	"github.com/tunx321/movie/internal/movies"
	transportHttp "github.com/tunx321/movie/internal/transport/http"
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

	httpHandler := transportHttp.NewHandler(mvService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}
	return nil
}

func main() {
	fmt.Println("Go REST API movie")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
