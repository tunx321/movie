package db

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func (d *Database) MigrateDB() error {
	fmt.Println("migrating out database")

	driver, err := postgres.WithInstance(d.Client, &postgres.Config{}) 
	if err != nil{
		return fmt.Errorf("could not connect the postgres driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres",
		driver,
	)

	if err != nil{
		fmt.Println(err)
		return err
	}

	if err := m.Up(); err != nil{
		if !errors.Is(err, migrate.ErrNoChange){
			return fmt.Errorf("couldn't run up migrations: %w", err)
		}
	}

	fmt.Println("successfully migrated the database")
	return nil
}
