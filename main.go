package main

import (
	"Go_Bookstore_Management/routes"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Starting server at port ", os.Getenv("PORT"), "\n")
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), r))
}
