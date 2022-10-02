package main

import (
	"github.com/gorilla/mux"
	h "kchoi85.io/rest-api/handlers"
	"kchoi85.io/rest-api/models"
	"log"
	"net/http"
)

func main() {
	model.InitDB()
	r := mux.NewRouter()

	r.HandleFunc("/api/books", h.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", h.GetBook).Methods("GET")
	r.HandleFunc("/api/books", h.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", h.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", h.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
