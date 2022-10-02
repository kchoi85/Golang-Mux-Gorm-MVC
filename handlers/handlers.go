package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	m "kchoi85.io/rest-api/models"
	"net/http"
)

// Init books var as a slice Book struct
var Books []m.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var books []m.Book
	m.DB.Find(&books)

	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var book m.Book
	m.DB.First(&book, params["id"])

	var author m.Author
	m.DB.First(&author, book.AuthorID)
	book.Author = author

	json.NewEncoder(w).Encode(book)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book m.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	m.DB.Create(&book)

	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var book m.Book
	m.DB.First(&book, params["id"])
	_ = json.NewDecoder(r.Body).Decode(&book)
	m.DB.Save(&book)

	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var book m.Book
	m.DB.Delete(&book, params["id"])

	json.NewEncoder(w).Encode(book)
}
