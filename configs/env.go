package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func EnvMessages() map[string]string {
	messages := make(map[string]string)
	messages["Fizz_MESSAGE"] = os.Getenv("Fizz_MESSAGE")
	messages["Buzz_MESSAGE"] = os.Getenv("Buzz_MESSAGE")
	messages["FizzBuzz_MESSAGE"] = os.Getenv("FizzBuzz_MESSAGE")

	return messages
}

func EnvPort() int {
	port, err := strconv.Atoi(os.Getenv("Port"))
	if err != nil {
		return 8080
	}
	return port
}
