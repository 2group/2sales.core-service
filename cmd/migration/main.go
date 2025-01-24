package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/2group/2sales.core-service/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	var migrationPath, migrationTable string
	flag.StringVar(&migrationPath, "migrations-path", "", "path to migrations")
	flag.StringVar(&migrationTable, "migration-table", "migrations", "name of migration table")
	cfg := config.MustLoad()
	flag.Parse()

	if migrationPath == "" {
		panic("migration path not defined")
	}

	postgresURL := cfg.Psql.Url

	m, err := migrate.New("file://"+migrationPath, postgresURL)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migration to apply")
			return
		}
		panic(err)
	}

	fmt.Println("migration applied successfully")
}
