package initializers

import (
	"Application/services"
	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	// LOAD .env FILE
	envError := godotenv.Load(".env")
	services.CheckError("initializers/env_initializers.go", 11, envError)
}
