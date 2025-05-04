package repository

import (
	"devSystem/models"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func NewPostgresDB(cfg *models.DevelopmentSystemDBConfig, password string) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Username, password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	return db, nil
}
func NewDatabase(config *models.ConfigService, env *models.Environment) *sqlx.DB {
	fmt.Println("starting database connection...")

	db, err := NewPostgresDB(&config.DevelopmentSystemDB, env.DBPassword)
	if err != nil {
		logrus.Fatalf("failed to initialize development database: %s", err.Error())
	}

	fmt.Println("database connected")
	return db
}
