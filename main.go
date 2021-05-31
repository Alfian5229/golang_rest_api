package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "alfian:laptopasus5229@tcp(127.0.0.1:3306)/golang")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	router := mux.NewRouter()
	http.ListenAndServe(":8000", router)
}
