package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	PORT     string
	DBName   string
	MongoURI string
}

var GlobalEnv Env

func init() {
	var ok bool
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Failed to detect .env file!")
	}

	GlobalEnv.PORT, ok = os.LookupEnv("PORT")
	if !ok {
		log.Fatal("No PORT detected!")
	}

	GlobalEnv.DBName, ok = os.LookupEnv("DB_NAME")
	if !ok {
		log.Fatal("No DBName detected!")
	}

	GlobalEnv.MongoURI, ok = os.LookupEnv("MONGO_URI")
	if !ok {
		log.Fatal("No MONGO_URI detected!")
	}
}
