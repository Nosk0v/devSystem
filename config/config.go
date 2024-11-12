package config

import (
	"encoding/json"
	"log"
	"os"

	"devSystem/internal/repository"
)

func Config(path string) (*repository.Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Error loading config:", err)
		return nil, err
	}

	var config repository.Config
	err = json.Unmarshal(b, &config)
	if err != nil {
		log.Fatal("Error parsing config:", err)
		return nil, err
	}

	return &config, nil
}
