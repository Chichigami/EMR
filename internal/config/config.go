package config

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/chichigami/EMR/internal/database"
	"github.com/chichigami/EMR/internal/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func LoadConfig() (*Config, *sql.DB) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}
	dbConnection, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}
	if err := dbConnection.Ping(); err != nil {
		log.Fatalf("Error pinging database: %s", err)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET must be set")
	}

	dbQueries := database.New(dbConnection)

	cfg := &Config{
		Database:  dbQueries,
		JWTSecret: jwtSecret,
	}
	return cfg, dbConnection
}

type Config struct {
	Database       *database.Queries
	JWTSecret      string
	DashboardCache map[time.Time]string
	PatientCache   map[string]models.Patient
}
