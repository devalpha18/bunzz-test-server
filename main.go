package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/devalpha18/bunzz-test-server/configs"
	"github.com/devalpha18/bunzz-test-server/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Setting up server, enabling CORS . . .")
	configs.LoadEnv()

	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"GET, POST, PUT, DELETE, OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router := mux.NewRouter()
	// routes
	routes.FizzBuzzRoute(router)

	port := configs.EnvPort()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handlers.CORS(credentials, methods, origins)(router)))
}
