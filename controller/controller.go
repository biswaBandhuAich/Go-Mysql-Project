package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	bookService "github.com/biswaBandhuAich/mysql/dal"
	"github.com/biswaBandhuAich/mysql/model"
	"github.com/gorilla/mux"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bookCollection := bookService.GetAllBooks()
	json.NewEncoder(w).Encode(&bookCollection)
}

// Fetch One Bok
func GetOneBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	bookId := params["id"]
	book := bookService.GetOneBook(bookId)
	json.NewEncoder(w).Encode(&book)
}

// Adding one Book
func AddOneBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book model.Book

	json.NewDecoder(r.Body).Decode(&book)

	bookService.AddOneBook(&book)

	fmt.Println("Book added")
	json.NewEncoder(w).Encode("Book added Succesfully")
}

// Adding many Books
func AddBooks(w http.ResponseWriter, r *http.Request) {
	//set header
	w.Header().Set("Content-Type", "application/json")
	booksToBeAdded := []model.Book{}
	json.NewDecoder(r.Body).Decode(&booksToBeAdded)

	for _, value := range booksToBeAdded {

		//call add one book
		bookService.AddOneBook(&value)

	}

}

// update Single Book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	book := model.Book{}
	json.NewDecoder(r.Body).Decode(&book)

	bookService.UpdateBook(&book)
}

// Delete one book
func DeleteOneBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	bookService.DeleteOneBook(params["id"])
	json.NewEncoder(w).Encode("Book Succesfully Delted")
}

func Home(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Welcome to the Books Library!!")
}
