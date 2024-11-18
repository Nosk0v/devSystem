// @title devSystem API  Documentation
// @version 1.0
// @description This is a sample server for an employee development system called "Competentum"
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
	"github.com/execaus/exloggo"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func main() {

	config, err := config.Config("../config/config.json")
	if err != nil {
		exloggo.Fatalf("failed to load configuration: %v", nil, err)
	}

	db, err := setupDatabase(config)
	if err != nil {
		exloggo.Fatalf("failed to connect to database: %v", nil, err)
	}
	defer db.Close()

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	usecase := usecase.NewUsecase(service)
	handler := handler.NewHandler(usecase)

	srv := server.Server{}
	runServer(&srv, handler, "8080")
	srv.Shutdown(db, context.Background())
}

func setupDatabase(config *repository.Config) (*sqlx.DB, error) {
	db, err := repository.NewPostgresConnection(config)
	if err != nil {
		return nil, err
	}
	exloggo.Info("database connection established successfully")
	return db, nil
}

func runServer(srv *server.Server, handler *handler.Handler, port string) {
	ginEngine := handler.InitRoutes()

	if err := srv.Run(port, ginEngine); err != nil {
		if err.Error() != "http: Server closed" {
			exloggo.Fatalf("error occurred while running http server: %s", nil, err.Error())
		}
	}
}
