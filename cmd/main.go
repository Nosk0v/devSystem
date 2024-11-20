// @title devSystem API Документация
// @version 1.0
// @description Это пример API сервера для системы развития сотрудников под названием «Компетентум».
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email alexandernoskov.dev@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /

package main

import (
	"context"
	"devSystem/config"
	"devSystem/internal/handler"
	"devSystem/internal/repository"
	"devSystem/internal/service"
	"devSystem/internal/usecase"
	"devSystem/server"
	"fmt"
	"github.com/execaus/exloggo"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"os"
)

func main() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "../config/config.json" // Умолчательный путь
	}

	config, err := config.Config(configPath)
	if err != nil {
		exloggo.Fatalf("failed to load configuration: %v", err)
	}

	db, err := setupDatabase(config)
	if err != nil {
		exloggo.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	skipMigrations := os.Getenv("SKIP_MIGRATIONS")
	if skipMigrations == "true" {
		exloggo.Info("Skipping database migrations as per configuration")
	} else {
		if err := applyMigrations(db); err != nil {
			exloggo.Fatalf("failed to apply migrations: %v", err)
		}
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	usecase := usecase.NewUsecase(service)
	handler := handler.NewHandler(usecase)

	srv := server.Server{}
	runServer(&srv, handler, "8080")

	srv.Shutdown(db, context.Background())
}

func setupDatabase(config *repository.Config) (*sqlx.DB, error) {
	dbURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.Password, config.DBName, config.SSLMode)

	db, err := sqlx.Connect("pgx", dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	exloggo.Info("database connection established successfully")
	return db, nil
}

func applyMigrations(db *sqlx.DB) error {
	migrationsDir := "./db/migrations" // менять в случае необходимости
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("failed to set dialect for migrations: %w", err)
	}
	if err := goose.Up(db.DB, migrationsDir); err != nil {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	exloggo.Info("migrations applied successfully")
	return nil
}

func runServer(srv *server.Server, handler *handler.Handler, port string) {
	ginEngine := handler.InitRoutes()

	if err := srv.Run(port, ginEngine); err != nil {
		if err.Error() != "http: Server closed" {
			exloggo.Fatalf("error occurred while running http server: %s", err.Error())
		}
	}
}
