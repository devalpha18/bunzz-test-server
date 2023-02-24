package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMessages() map[string]string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	messages := make(map[string]string)
	messages["Fizz_MESSAGE"] = os.Getenv("Fizz_MESSAGE")
	messages["Buzz_MESSAGE"] = os.Getenv("Buzz_MESSAGE")
	messages["FizzBuzz_MESSAGE"] = os.Getenv("FizzBuzz_MESSAGE")

	return messages
}
