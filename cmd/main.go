// @title devSystem API Документация
// @version 1.0
// @description Это API сервер для сервиса «Развитие сотрудников» с целью автоматизации процессов управления персоналом.
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
	"devSystem/internal/handler"
	"devSystem/internal/repository"
	"devSystem/internal/service"
	"devSystem/internal/usecase"
	"devSystem/models"
	"devSystem/server"
	"encoding/json"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file: %s", err.Error())
	}

	environment := &models.Environment{}
	configService := &models.ConfigService{}

	if err := runLogger(); err != nil {
		logrus.Fatal(err.Error())
	}

	if err := loadEnvironment(environment); err != nil {
		logrus.Fatalf(err.Error())
		return
	}
	if err := loadConfig(configService); err != nil {
		logrus.Fatalf(err.Error())
		return
	}
	logrus.Info("load local config success")

	logrus.Infof("Server UTC time: %v", time.Now().UTC())

	db := repository.NewDatabase(configService, environment)
	defer func() {
		if err := db.Close(); err != nil {
			logrus.Errorf("failed to close db connection: %v", err)
		}
	}()

	skipMigrations := os.Getenv("SKIP_MIGRATIONS")
	if skipMigrations == "true" {
		logrus.Info("Skipping database migrations as per configuration")
	} else {
		if err := applyMigrations(db); err != nil {
			logrus.Fatalf("failed to apply migrations: %v", err)
		}
	}

	repositorySources := repository.Sources{
		Db: db,
	}

	repo := repository.NewRepository(&repositorySources)
	service := service.NewService(repo, configService)
	usecase := usecase.NewUsecase(service)
	handler := handler.NewHandler(usecase)

	srv := server.Server{}
	runServer(&srv, handler, configService.Server.Port)

	srv.Shutdown(db, context.Background())
}

func applyMigrations(db *sqlx.DB) error {
	migrationsDir := "./db/migrations"
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("failed to set dialect for migrations: %w", err)
	}
	if err := goose.Up(db.DB, migrationsDir); err != nil {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}
	logrus.Info("migrations applied successfully")
	return nil
}

func runServer(srv *server.Server, handler *handler.Handler, port string) {
	logrus.Infof("Server started at UTC time: %v", time.Now().UTC())
	ginEngine := handler.InitRoutes()

	if err := srv.Run(port, ginEngine); err != nil {
		if err.Error() != "http: Server closed" {
			logrus.Fatalf("error occurred while running http server: %s", err.Error())
		}
	}
}

func runLogger() error {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&CustomFormatter{})

	currentTime := time.Now()
	yearMonthDir := fmt.Sprintf("logs/%d-%02d", currentTime.Year(), currentTime.Month())

	err := os.MkdirAll(yearMonthDir, os.ModePerm)
	if err != nil {
		logrus.Error(err.Error())
		return err
	}

	logFile := fmt.Sprintf("%s/%d-%02d-%02d.log", yearMonthDir, currentTime.Year(), currentTime.Month(), currentTime.Day())

	logFileHandle, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Error(err.Error())
		return err
	}

	logrus.SetOutput(io.MultiWriter(os.Stdout, logFileHandle))

	return nil
}

type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	prefixPath := pwd + "/"

	shortFilePath := strings.TrimPrefix(filepath.ToSlash(entry.Caller.File), filepath.ToSlash(prefixPath))

	var fields string
	for key, value := range entry.Data {
		fields += fmt.Sprintf("\"%s\":\"%v\",", key, value)
	}

	if len(fields) > 0 {
		fields = fields[:len(fields)-1]
	}

	if len(fields) > 0 {
		fields = ", " + fields
	}

	log := fmt.Sprintf(
		"{\"level\":\"%s\",\"msg\":\"%s\",\"point\": \" %s:%d \",\"short_point\":\"%s:%d\", \"time\":\"%s\"%s}\n",
		entry.Level.String(),
		entry.Message,
		entry.Caller.File,
		entry.Caller.Line,
		shortFilePath,
		entry.Caller.Line,
		entry.Time.Format(time.RFC3339),
		fields,
	)
	return []byte(log), nil
}

func loadEnvironment(environment *models.Environment) error {
	if err := godotenv.Load(".env"); err != nil {
		logrus.Warning("load file not found, Environment variables load from Environment")
	}
	if err := env.Parse(environment); err != nil {
		return err
	}

	return nil
}

func loadConfig(config *models.ConfigService) error {
	file, err := os.ReadFile("./config/config.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return err
	}

	return nil
}
