package routes

import (
	"github.com/devalpha18/bunzz-test-server/controllers"
	"github.com/gorilla/mux"
)

func FizzBuzzRoute(router *mux.Router) {
	router.HandleFunc("/fizzbuzz", controllers.GetFizzBuzz()).Methods("POST")
}
