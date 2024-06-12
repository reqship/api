package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVarByName(name string) string {
	err := godotenv.Load()

	if err != nil {
		env_file := ".env"
		if os.Args[0] == "./main" {
			env_file = ".example.env"
		}
		err = godotenv.Load(env_file)

		if err != nil {
			log.Fatal("error loading .env file")
		}
	}

	return os.Getenv(name)
}
