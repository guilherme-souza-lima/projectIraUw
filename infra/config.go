package infra

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Environment      string
	App              string
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
}

func NewConfig() Config {
	if os.Getenv("ENVIRONMENT") == "" {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatalln("Error loading env file")
		}
	}

	return Config{
		Environment:      os.Getenv("APP"),
		App:              os.Getenv("ENVIRONMENT"),
		PostgresHost:     os.Getenv("HOST_POSTGRES"),
		PostgresPort:     os.Getenv("PORT_POSTGRES"),
		PostgresUser:     os.Getenv("USER_POSTGRES"),
		PostgresPassword: os.Getenv("PASSWORD_POSTGRES"),
		PostgresDatabase: os.Getenv("DATABASE_POSTGRES"),
	}

}
