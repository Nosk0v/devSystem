// cmd/main.go
package main

import (
	"devSystem/config"
	"devSystem/internal/repository"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

func main() {
	// Инициализация базы данных
	db := config.InitDB()
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)

	// Создание общего репозитория
	repo := repository.NewRepository(db)

	materials, err := repo.GetAllMaterials()
	if err != nil {
		log.Fatalf("Ошибка получения материалов: %v", err)
	}
	fmt.Println("Материалы:", materials)

	competencies, err := repo.GetAllCompetencies()
	if err != nil {
		log.Fatalf("Ошибка получения компетенций: %v", err)
	}
	fmt.Println("Компетенции:", competencies)
}
