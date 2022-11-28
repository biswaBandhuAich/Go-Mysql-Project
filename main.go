package main

import (
	"fmt"
	"net/http"

	"github.com/biswaBandhuAich/mysql/router"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Welcome to MYSQL - GO ")

	fmt.Println("Connection Established Succesfull")
	r := router.Router()

	//serve
	http.ListenAndServe(":4000", r)

}
