package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver = "mysql"
	dbUser   = "root"
	dbPass   = "pass"
	dbName   = "form-go-db"
)

func main() {

	http.HandleFunc("/", indexPageHandler)
	http.HandleFunc("/submit", userFormSubmitHandler)

	fmt.Println("listening at localhost port 8000")

	http.ListenAndServe(":8000", nil)
}

func indexPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "invalid request to the route", http.StatusInternalServerError)
		return
	}
	templ := template.Must(template.ParseFiles("index.html"))
	templ.Execute(w, nil)
}

func userFormSubmitHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello  sir")
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if r.Method != http.MethodPost {
		templ := template.Must(template.ParseFiles("index.html"))
		templ.Execute(w, nil)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	query := "insert into users (name, email) values (?, ?)"
	_, err = db.Exec(query, name, email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "User added")

}
