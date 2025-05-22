package models

type Environment struct {
	DBUser       string `env:"DB_USER"`
	DBPassword   string `env:"DB_PASSWORD"`
	DBHost       string `env:"DB_HOST"`
	DBPort       string `env:"DB_PORT"`
	DBName       string `env:"DB_NAME"`
	DBSSLMode    string `env:"DB_SSLMODE"`
	JWTSecretKey string `env:"JWT_SECRET_KEY"`
}
