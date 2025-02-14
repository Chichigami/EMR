package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/chichigami/EMR/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

func LoadConfig() (*Config, func()) {
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

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})
	dbQueries := database.New(dbConnection)

	cfg := &Config{
		Database:  dbQueries,
		JWTSecret: jwtSecret,
		Cache:     rdb,
	}

	cleanup := func() {
		if err := dbConnection.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
		if err := rdb.Close(); err != nil {
			log.Printf("Error closing redis: %v", err)
		}
	}

	return cfg, cleanup
}

type Config struct {
	Database  *database.Queries
	JWTSecret string
	Cache     *redis.Client
}
