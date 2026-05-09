package migrations

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	mpostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	_ "github.com/lib/pq"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

func Migrate(dsn string, targetVersion uint) error {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("open migration db: %w", err)
	}
	defer db.Close()

	driver, err := mpostgres.WithInstance(db, &mpostgres.Config{})
	if err != nil {
		return fmt.Errorf("create migration driver: %w", err)
	}

	sourceDriver, err := iofs.New(migrationFiles, "migrations")
	if err != nil {
		return fmt.Errorf("create migration source: %w", err)
	}

	m, err := migrate.NewWithInstance("iofs", sourceDriver, "postgres", driver)
	if err != nil {
		return fmt.Errorf("create migration instance: %w", err)
	}
	defer func() {
		srcErr, dbErr := m.Close()
		if srcErr != nil {
			fmt.Printf("close migration source error: %v\n", srcErr)
		}
		if dbErr != nil {
			fmt.Printf("close migration db error: %v\n", dbErr)
		}
	}()

	currentVersion, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		return fmt.Errorf("read migration version: %w", err)
	}

	if dirty {
		return fmt.Errorf("database is in dirty migration state")
	}

	if err == migrate.ErrNilVersion {
		currentVersion = 0
	}

	step := int(targetVersion) - int(currentVersion)
	if step == 0 {
		return nil
	}

	if err := m.Steps(step); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("apply migration steps: %w", err)
	}

	return nil
}
