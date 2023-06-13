package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // postgres driver

	"github.com/kokoichi206/awesome-chat-app/backend/repository"
	"github.com/kokoichi206/awesome-chat-app/backend/util/logger"
)

type database struct {
	db     *sql.DB
	logger logger.Logger
}

func Connect(driver, host, port, user, password, dbname, sslmode string) (*sql.DB, error) {
	source := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)

	sqlDB, err := sql.Open(driver, source)
	if err != nil {
		return nil, fmt.Errorf("failed to open sql: %w", err)
	}

	return sqlDB, nil
}

func New(sqlDB *sql.DB, logger logger.Logger) repository.Database {
	db := &database{
		db:     sqlDB,
		logger: logger,
	}

	return db
}
