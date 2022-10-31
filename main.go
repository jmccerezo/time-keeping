package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/jmccerezo/time-keeping/views"
)

func main() {
	fmt.Println("Welcome To Simple Time Keeping CRUD App")
	fmt.Println("Go to http://localhost:1234/")

	CreateDB("time_keeping")
	CreateTable("users", "time_keeping")
	Handlers()

	http.ListenAndServe(":1234", nil)
}

func Handlers() {
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates/"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", views.LogInHander)
	http.HandleFunc("/dashboard/", views.DashboardHandler)
	http.HandleFunc("/delete/", views.DeleteHandler)
	http.HandleFunc("/logout/", views.LogoutHandler)
	http.HandleFunc("/signup/", views.SignUpHandler)
	http.HandleFunc("/update/", views.UpdateHandler)
}

func CreateDB(name string) *sql.DB {
	db, err := sql.Open("mysql", "root:Allen is Great 200%@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + name)
	if err != nil {
		panic(err)
	}
	db.Close()

	db, err = sql.Open("mysql", "root:Allen is Great 200%@tcp(127.0.0.1:3306)/"+name)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return db
}

func CreateTable(table string, name string) *sql.DB {
	db, err := sql.Open("mysql", "root:Allen is Great 200%@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("USE " + name)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS " + table + "(id INT(11) PRIMARY KEY AUTO_INCREMENT, username varchar(32) UNIQUE, password varchar (32), first_name varchar (32), last_name varchar (32), time_in datetime, time_out datetime);")
	if err != nil {
		panic(err)
	}
	db.Close()

	return db
}
