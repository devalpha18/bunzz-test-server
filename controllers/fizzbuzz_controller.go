package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/devalpha18/bunzz-test-server/configs"
	"github.com/devalpha18/bunzz-test-server/entities"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func GetFizzBuzz() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var fizzbuzz entities.FizzBuzzRequest

		// validate the request body
		if err := json.NewDecoder(r.Body).Decode(&fizzbuzz); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := entities.FizzBuzzResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		// use the validator library to validate required fields
		if validationErr := validate.Struct(&fizzbuzz); validationErr != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := entities.FizzBuzzResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		result := calcFizzBuzz(fizzbuzz.Count)

		rw.WriteHeader(http.StatusOK)
		response := entities.FizzBuzzResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}}
		json.NewEncoder(rw).Encode(response)
	}
}

func calcFizzBuzz(input uint) string {
	if input%15 == 0 {
		return configs.EnvMessages()["FizzBuzz_MESSAGE"]
	}
	if input%3 == 0 {
		return configs.EnvMessages()["Fizz_MESSAGE"]
	}
	if input%5 == 0 {
		return configs.EnvMessages()["Buzz_MESSAGE"]
	}
	return ""
}
