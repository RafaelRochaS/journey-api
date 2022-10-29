package utils

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var HOST string
var KAFKA_HOSTS string
var KAFKA_CLIENT_ID string
var KAFKA_TOPIC string
var TRUSTED_PROXIES string
var PORT int

func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	HOST = os.Getenv("HOST")
	KAFKA_HOSTS = os.Getenv("KAFKA_HOSTS")
	KAFKA_CLIENT_ID = os.Getenv("KAFKA_CLIENT_ID")
	KAFKA_TOPIC = os.Getenv("KAFKA_TOPIC")
	TRUSTED_PROXIES = os.Getenv("TRUSTED_PROXIES")
	PORT, err = strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		PORT = 8080
	}
}
