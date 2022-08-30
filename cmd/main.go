package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"

	"github.com/rogaliiik/bookstore/pkg/routes"
)

// Loads values from .env into the system
func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

// Server processing
func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)

	http.Handle("/", router)
	fmt.Println("Server is working at localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
