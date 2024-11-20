package repository

import (
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib" // Импортируем драйвер для PostgreSQL
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"DBName"`
	SSLMode  string `json:"SSLMode"`
}

func NewPostgresConnection(config *Config) (*sqlx.DB, error) {
	// Формируем строку подключения
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.Password, config.DBName, config.SSLMode)

	// Подключаемся к базе данных
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	return db, nil
}
