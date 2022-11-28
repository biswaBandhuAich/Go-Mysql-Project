package bookService

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/biswaBandhuAich/mysql/model"
)

func GetAllBooks() *[]model.Book {

	var books []model.Book

	dbCon := getDbConnection()
	defer dbCon.Close()

	result, err := dbCon.Query("select * from books")

	if err != nil {
		fmt.Println("Error occured while fetching all books")
	}

	for result.Next() {
		var book model.Book

		err := result.Scan(&book.ID, &book.Name, &book.Price, &book.Genre, &book.AuthorID)
		fmt.Println("The book is", book)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}
	fmt.Println("All book is", books)

	return &books
}

// Fetch One Bok
func GetOneBook(bookId string) model.Book {
	dbCon := getDbConnection()
	defer dbCon.Close()

	result, err := dbCon.Query("Select * from books where id=?", bookId)

	if err != nil {
		log.Fatal("Error occured while fetching Book")
	}

	var book model.Book

	if result.Next() {

		err2 := result.Scan(&book.ID, &book.Name, &book.Price, &book.Genre, &book.AuthorID)

		if err2 != nil {
			log.Fatal("Error occured while parsing data")
		}
	}
	return book
}

// Adding one Book
func AddOneBook(book *model.Book) {

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
}

func DeleteOneBook(bookId string) {
	dbCon := getDbConnection()
	defer dbCon.Close()
	sql := "delete from books where id=?"
	_, err := dbCon.Exec(sql, bookId)
	if err != nil {
		log.Fatal("Error occured while deleting : ", bookId)
	}
}

func DeleteAll() {
	dbCon := getDbConnection()
	defer dbCon.Close()
	sql := "delete from books"
	_, err := dbCon.Exec(sql)
	if err != nil {
		log.Fatal("Error occured while deleting all books ")
	}
}

func UpdateBook(book *model.Book) {
	dbCon := getDbConnection()
	defer dbCon.Close()

	id := book.ID

	DeleteOneBook(id)
	AddOneBook(book)
}

func getDbConnection() *sql.DB {

	dbCon, err := sql.Open("mysql", "root:root123456@tcp(127.0.0.1:3306)/bookzone")

	if err != nil {
		log.Fatal(err)
	}

	return dbCon
}
