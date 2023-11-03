package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Book struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
	ID          string `json:"ID"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, item := range books {
		if item.ID == params["ID"] {

			books = append(books[:i], books[i+1:]...)
		}

	}
	json.NewEncoder(w).Encode(books)

}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["ID"] {
			json.NewEncoder(w).Encode(item)

		}

	}

}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(books)

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range books {
		if item.ID == params["ID"] {
			books = append(books[:i], books[i+1:]...)
		}
	}

	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	book.ID = params["ID"]
	books = append(books, book)
	json.NewEncoder(w).Encode(books)

}

func main() {
	r := mux.NewRouter()
	logger := logrus.New()
	books = append(books, Book{
		Name:        "example",
		Author:      "example",
		Publication: "example",
		ID:          "1",
	})
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{ID}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{ID}", updateBook).Methods("POST")
	r.HandleFunc("/books/{ID}", deleteBook).Methods("DELETE")
	logger.Fatal(http.ListenAndServe(":8080", r))
}
