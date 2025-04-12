package config

import (
	"log"
	"os"
)

func MustGetenv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Environment variable %s is not set!", key)
	}
	return val
}

func GetenvWithDefault(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}

var DATABASE_URL = MustGetenv("DATABASE_URL")
var PORT = GetenvWithDefault("PORT", "3000")
var LISTEN_ON = ":" + PORT
var APP_ENV = GetenvWithDefault("APP_ENV", "dev")

var LOG_TIMESTAMPS = APP_ENV == "dev"
