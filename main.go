package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to MYSQL - GO ")

	fmt.Println("Connection Established Succesfull")

	//Initialize router
	r := mux.NewRouter()

	//handler
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/book/{id}", getOneBook).Methods("GET")
	r.HandleFunc("/books", getAllBooks).Methods("GET")
	r.HandleFunc("/book", addOneBook).Methods("POST")
	r.HandleFunc("/books", addBooks).Methods("POST")

	//serve
	http.ListenAndServe(":4000", r)

}

// Fetch all books
func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var books []Book

	dbCon := getDbConnection()
	defer dbCon.Close()

	result, err := dbCon.Query("select * from books")

	if err != nil {
		fmt.Println("Error occured while fetching all books")
	}

	for result.Next() {
		var book Book

		err := result.Scan(&book.ID, &book.Name, &book.Price, &book.Genre, &book.AuthorID)
		fmt.Println("The book is", book)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}
	fmt.Println("All book is", books)

	json.NewEncoder(w).Encode(&books)
}

// Fetch One Bok
func getOneBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	bookId := params["id"]

	dbCon := getDbConnection()
	defer dbCon.Close()

	result, err := dbCon.Query("Select * from books where id=?", bookId)

	if err != nil {
		log.Fatal("Error occured while fetching Book")
	}

	if result.Next() {
		var book Book

		err2 := result.Scan(&book.ID, &book.Name, &book.Price, &book.Genre, &book.AuthorID)

		if err2 != nil {
			log.Fatal("Error occured while parsing data")
		}
		json.NewEncoder(w).Encode(book)
	} else {
		json.NewEncoder(w).Encode("Book not found")
	}

}

// Adding one Book
func addOneBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book

	json.NewDecoder(r.Body).Decode(&book)

	dbCon := getDbConnection()
	defer dbCon.Close()

	sql := "INSERT INTO books values(?,?,?,?,?)"

	authorId, _ := strconv.Atoi(book.AuthorID)
	price, _ := strconv.Atoi(book.Price)
	bookId, _ := strconv.Atoi(book.ID)

	result, err := dbCon.Exec(sql, bookId, book.Name, price, book.Genre, authorId)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Book added", result)
	json.NewEncoder(w).Encode("Book added Succesfully")
}

// Adding many Books
func addBooks(w http.ResponseWriter, r *http.Request) {
	//set header
	w.Header().Set("Content-Type", "application/json")
	booksToBeAdded := []Book{}
	json.NewDecoder(r.Body).Decode(&booksToBeAdded)

	for value := range booksToBeAdded {

		//call add one book

	}

}

// Update multiple Books
func updateBooks(w http.ResponseWriter, r *http.Request) {

}

// update Single Book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

func home(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Welcome to the Books Library!!")
}

func getDbConnection() *sql.DB {

	dbCon, err := sql.Open("mysql", "root:root123456@tcp(127.0.0.1:3306)/bookzone")

	if err != nil {
		log.Fatal(err)
	}

	return dbCon
}
