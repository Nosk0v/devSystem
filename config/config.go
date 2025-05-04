package config

import (
	"devSystem/models"
	"encoding/json"
	"log"
	"os"
)

func Config(path string) (*models.ConfigService, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Error loading config:", err)
		return nil, err
	}

	var config models.ConfigService
	err = json.Unmarshal(b, &config)
	if err != nil {
		log.Fatal("Error parsing config:", err)
		return nil, err
	}

	return &config, nil
}
