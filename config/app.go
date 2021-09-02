package config

import (
	"log"

	"github.com/joho/godotenv"
)

var myEnv map[string]string

func init() {
	tmpEnv, err := godotenv.Read()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	myEnv = tmpEnv
}

func Param(param_name string) string {
	value, ok := myEnv[param_name]

	if ok {
		return value
	}

	return ""
}
