package migrations

import (
	"database/sql"
	"fmt"

	"crud_api_go/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Migration(config *config.PostgreSQL, version int) error {
	connection := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		config.USER,
		config.PASSWORD,
		config.HOST,
		config.PORT,
		config.DATABASE_NAME)

	db, err := sql.Open("postgres", connection)
	defer db.Close()
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		return err
	}

	if err := m.Steps(version); err != nil {
		return err
	}

	defer m.Close()
	return nil
}
