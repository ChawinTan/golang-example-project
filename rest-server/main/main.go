package main

import (
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
	"server/services"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "chawin:@password@tcp(127.0.0.1:3306)/golang")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	fmt.Println("----- database started successfully -----")

	router := mux.NewRouter()

	router.HandleFunc("/get_user", func (w http.ResponseWriter, r *http.Request) {
		services.GetUser(w, r, db)
	}).Methods("GET")

	http.ListenAndServe(":8002", router)
}