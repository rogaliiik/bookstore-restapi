package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rogaliiik/bookstore/pkg/config"
	"github.com/rogaliiik/bookstore/pkg/routes"
)

// Server processing
func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)

	http.Handle("/", router)
	address := fmt.Sprintf(":%s", config.Conf.Port)
	fmt.Println("Server is working at port:", config.Conf.Port)
	log.Fatal(http.ListenAndServe(address, router))
}
