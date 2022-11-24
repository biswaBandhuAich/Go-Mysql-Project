package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

	//serve
	http.ListenAndServe(":4000", r)

}

// Fetch all books
func getAllBooks(w http.ResponseWriter, r *http.Request) {

	dbCon := getDbConnection()
	defer dbCon.Close()

	dbCon.Query("select * from books")

}

// Fetch One Bok
func getOneBook(w http.ResponseWriter, r *http.Request) {

}

// Adding one Book
func addOneBook(w http.ResponseWriter, r *http.Request) {

}

// Adding many Books
func addBooks(w http.ResponseWriter, r *http.Request) {

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
