package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoUri() string {
	// cargar el env
	err := godotenv.Load()
	// handler err
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
	// retorno
	return os.Getenv("MONGOURI")
}
