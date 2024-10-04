package database

import (
	"context"
	"fmt"

	"example.com/m/ent"
	"example.com/m/internal/config"
	_ "github.com/lib/pq"
)

type Database struct {
	Client *ent.Client
}

var DB *Database = nil

func New(cfg *config.Config) (*Database, error) {
	postgresConnection := fmt.Sprintf(
		"host=%v port=%v user=%v dbname=%v password=%v sslmode=%v",
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbUser,
		cfg.DbName,
		cfg.DbPassword,
		cfg.DbSSLMode,
	)

	client, err := ent.Open("postgres", postgresConnection)
	if err != nil {
		return &Database{}, fmt.Errorf("failed opening connection to postgres: %w", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		return &Database{}, fmt.Errorf("failed to create schema resources: %w", err)
	}

	DB = &Database{
		Client: client,
	}

	return DB, nil
}

func (db *Database) Close() error {
	if err := db.Client.Close(); err != nil {
		return fmt.Errorf("database error: %w", err)
	}

	return nil
}
