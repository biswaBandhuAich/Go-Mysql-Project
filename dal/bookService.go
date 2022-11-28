package bookService

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/biswaBandhuAich/mysql/errorHandler"
	"github.com/biswaBandhuAich/mysql/model"
)

func GetAllBooks() *[]model.Book {

	var books []model.Book

	dbCon := getDbConnection()
	defer dbCon.Close()

	result, err := dbCon.Query("select * from books")
	errorHandler.Handle(err, "Error Occured while fetching all books")
	for result.Next() {
		var book model.Book

		err := result.Scan(&book.ID, &book.Name, &book.Price, &book.Genre, &book.AuthorID)
		errorHandler.Handle(err, "Error occured while parcing all books.")

		fmt.Println("The book is", book)
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

	errorHandler.Handle(err, "Error occured while fetching Book")

	var book model.Book

	if result.Next() {
		err2 := result.Scan(&book.ID, &book.Name, &book.Price, &book.Genre, &book.AuthorID)
		errorHandler.Handle(err2, "Error parsig object at : Get One Book")
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

	errorHandler.Handle(err, "Failed inserting into database")
	fmt.Println("Book added", result)
}

// Deleting single book
func DeleteOneBook(bookId string) {
	dbCon := getDbConnection()
	defer dbCon.Close()

	sql := "delete from books where id=?"

	_, err := dbCon.Exec(sql, bookId)
	errorHandler.Handle(err, "Error occured while deleting : "+bookId)
}

func DeleteAll() {
	dbCon := getDbConnection()
	defer dbCon.Close()
	sql := "delete from books"
	_, err := dbCon.Exec(sql)
	errorHandler.Handle(err, "Deleting failed")
}

func UpdateBook(book *model.Book) {
	dbCon := getDbConnection()
	defer dbCon.Close()

	id := book.ID

	// Removing the previous entry and replacing with same ID
	DeleteOneBook(id)
	AddOneBook(book)
}

func getDbConnection() *sql.DB {

	dbCon, err := sql.Open("mysql", "root:root123456@tcp(127.0.0.1:3306)/bookzone")

	errorHandler.Handle(err, "Unable to establish database connection")

	return dbCon
}
