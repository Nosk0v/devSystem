package models

type Environment struct {
	DBPassword   string `env:"DB_PASSWORD"`
	JWTSecretKey string `json:"jwt_secret_key" binding:"requires"`
}
