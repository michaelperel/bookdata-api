package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matt-FFFFFF/bookdata-api/datastore"
)

var (
	books datastore.BookStore
)

func init() {
	books = &datastore.Books{}
	books.Initialize()
}

func main() {
	r := mux.NewRouter()
	log.Println("bookdata api")
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", timer(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "api v1")
	}))
	api.HandleFunc("/books", timer(getAllBooks)).Methods(http.MethodGet)
	api.HandleFunc("/books/authors/{author}", timer(getAllBooksByAuthor)).Methods(http.MethodGet)
	api.HandleFunc("/books/title/{title}", timer(getAllBooksByTitle)).Methods(http.MethodGet)
	api.HandleFunc("/book/isbn/{isbn}", timer(getBookByISBN)).Methods(http.MethodGet)
	api.HandleFunc("/book/isbn/{isbn}", timer(deleteBookByISBN)).Methods(http.MethodDelete)
	api.HandleFunc("/book", timer(createBook)).Methods(http.MethodPost)
	log.Fatalln(http.ListenAndServe(":8080", r))
}
