package main

import (
	"context"
	"fmt"

	"github.com/tunx321/movie/internal/db"
)

func Run() error {
	fmt.Println("starting up our appplication")
	db, err := db.NewDatabase()
	if err != nil{
		fmt.Println("failed to ping to database")
		return err
	}

	if err := db.Ping(context.Background()); err != nil{
		return err
	}

	fmt.Println("successfully connected and pinged database")
	return nil
}

func main() {
	fmt.Println("Go REST API movie")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
