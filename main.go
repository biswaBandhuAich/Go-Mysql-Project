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

	dbCon, err := sql.Open("mysql", "root:root123456@tcp(127.0.0.1:3306)/bookzone")

	if err != nil {
		log.Fatal(err)
	}

	defer dbCon.Close()

	fmt.Println("Connection Established Succesfull")

	//Initialize router
	r := mux.NewRouter()

	//handler

	r.HandleFunc("/", home)

	//serve
	http.ListenAndServe(":4000", r)

}

// Fetch all books
func getAllBooks() {

}

// Fetch One Bok
func getOneBook() {

}

// Adding one Book
func addOneBook() {

}

// Adding many Books
func addBooks() {

}

// Update multiple Books
func updateBooks() {

}

// update Single Book
func updateBook() {

}

func home(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Welcome to the Books Library!!")
}
