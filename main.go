package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

// Book struct (Model)
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book
var nextID int = 1

// Create a new book
func createBook(w http.ResponseWriter, r *http.Request) {
	var newBook Book
	_ = json.NewDecoder(r.Body).Decode(&newBook)
	newBook.ID = nextID
	nextID++
	books = append(books, newBook)
	json.NewEncoder(w).Encode(newBook)
}

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

// Get a single book
func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for _, book := range books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.NotFound(w, r)
}

// Update a book
func updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, book := range books {
		if book.ID == id {
			_ = json.NewDecoder(r.Body).Decode(&books[i])
			books[i].ID = id // Preserve the ID
			json.NewEncoder(w).Encode(books[i])
			return
		}
	}
	http.NotFound(w, r)
}

// Delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			fmt.Fprintf(w, "Book with ID %d deleted", id)
			return
		}
	}
	http.NotFound(w, r)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", r)
}
