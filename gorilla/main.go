package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var books = []string{"Citas del Presidente Mao Tse-Tung", "Harry Potter", "El Señor de los Anillos", "El Alquimista"}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/teas/{category}/", TeasCategoryHandler).Methods("GET", "HEAD")
	r.HandleFunc("/teas/{category}/{id:[0-9]+}", TeaHandler)

	r.HandleFunc("/secure", SecureHandler).Schemes("https")
	r.HandleFunc("/insecure", InsecureHandler).Schemes("http")

	bookRouter := r.PathPrefix("/books").Subrouter()
	bookRouter.HandleFunc("/", AllBooks)
	bookRouter.HandleFunc("/{title}", GetBook)


	http.ListenAndServe(":8085", r)
}

func TeasCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

func TeaHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ID: %v\n", vars["id"])
}

func AllBooks(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Libros: %v\n", books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	book := ""

	for _, element := range books {
		if element == vars["title"] {
			book = element
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Libro: %v\n", book)
}

func SecureHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Modo: %v\n", "Seguro")
}

func InsecureHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Modo: %v\n", "Inseguro")
}