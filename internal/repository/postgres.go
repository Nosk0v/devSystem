package repository

import (
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"DBName"`
	SSLMode  string `json:"SSLMode"`
}
