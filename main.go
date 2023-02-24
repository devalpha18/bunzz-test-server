package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/devalpha18/bunzz-test-server/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Settin up server, enabling CORS . . .")

	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"GET, POST, PUT, DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router := mux.NewRouter()
	// routes
	routes.FizzBuzzRoute(router)
	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(credentials, methods, origins)(router)))
}
