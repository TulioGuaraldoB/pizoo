package env

import (
	"log"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

type EnvironmentVariable struct {
	PostgresUser     string
	PostgresPassword string
	PostgresHost     string
	PostgresPort     string
	PostgresDatabase string
}

var Env *EnvironmentVariable

func GetEnvironmentVariable() *EnvironmentVariable {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Failed to load .env file. %s", err.Error())
	}

	Env = &EnvironmentVariable{
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		PostgresDatabase: os.Getenv("POSTGRES_DATABASE"),
	}

	checkVariables()

	return Env
}

func checkVariables() {
	value := reflect.ValueOf(*Env)
	typeOfValue := value.Type()
	for index := 0; index < value.NumField(); index++ {
		envValue := value.Field(index).Interface()
		if envValue == "" {
			envName := typeOfValue.Field(index).Name
			envName = transformEnvText(envName)
			log.Printf("ERROR | %s", envName)
		}
	}

}
