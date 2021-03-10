package main

import (
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
	"server/services"
)

func main() {
	db, err := sql.Open("mysql", "root:password@/golang")

	err = db.Ping()

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