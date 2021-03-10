package services

import (
	"encoding/json"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"net/http"
)

type User struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

func GetUser(w http.ResponseWriter, r* http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json")
	var users []User

	result, err := db.Query("SELECT * from user")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		var user User
		err := result.Scan(&user.Id, &user.Name)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}
