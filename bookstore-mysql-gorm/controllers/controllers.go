package controllers

import (
	"encoding/json"
	"github.com/Go-By-Example-Projects/bookstore-mysql-gorm/models"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	allbooks := models.GetAll()

	res, err := json.Marshal(&allbooks)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	mybook := &models.Book{}
	w.Header().Set("content-type", "application/json")
	res, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(res, &mybook)
	if err != nil {
		panic(err.Error())
	}
	mybook.CreateBook()

	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	mid := params["id"]
	ID, _ := strconv.ParseInt(mid, 0, 0)
	_, bookDetails := models.GetBookById(ID)

	res, _ := json.Marshal(bookDetails)

	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
