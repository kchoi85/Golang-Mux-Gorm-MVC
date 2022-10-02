package main

import (
	"github.com/gorilla/mux"
	h "kchoi85.io/rest-api/handlers"
	m "kchoi85.io/rest-api/models"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	// Mock data - @TODO: implement DB conn
	h.Books = append(h.Books, m.Book{
		ID:    "1",
		Isbn:  "1234",
		Title: "Book One",
		Author: &m.Author{
			Firstname: "John",
			Lastname:  "Doe",
		},
	})
	h.Books = append(h.Books, m.Book{
		ID:    "2",
		Isbn:  "5678",
		Title: "Book Two",
		Author: &m.Author{
			Firstname: "Doe",
			Lastname:  "John",
		},
	})

	r.HandleFunc("/api/books", h.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", h.GetBook).Methods("GET")
	r.HandleFunc("/api/books", h.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", h.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", h.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
