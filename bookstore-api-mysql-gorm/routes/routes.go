package routes

import (
	"github.com/Go-By-Example-Projects/bookstore-mysql-gorm/controllers"
	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {

	r.HandleFunc("/books", controllers.GetAllBooks).Methods("GET")
	r.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", controllers.GetBookById).Methods("GET")
}
