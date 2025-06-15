package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:yourpassword@tcp(127.0.0.1:3306)/user_auth")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler) *
		http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
