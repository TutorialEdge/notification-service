package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/TutorialEdge/ctxlog"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// MigrateDB - runs all migrations in the migrations
func MigrateDB(db *sql.DB) error {
	ctx := context.Background()
	log := ctxlog.New(
		ctxlog.WithJSONFormat(),
	)
	log.Info(ctx, "migrating our database")
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Error(ctx, err.Error())
		return fmt.Errorf("could not create the postgres driver: %w", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres", driver)
	if err != nil {
		log.Error(ctx, err.Error())
		return err
	}
	err = m.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Info(ctx, "no change on migration")
		} else {
			// log the error which is likely
			log.Error(ctx, err.Error())
			return err
		}
	}

	return nil
}
