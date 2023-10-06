package main

import (
	"github.com/Go-By-Example-Projects/bookstore-mysql-gorm/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.Routes(r)
	http.ListenAndServe(":8080", r)

}
