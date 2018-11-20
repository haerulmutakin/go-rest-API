package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restapi/conn"

	"github.com/gorilla/mux"
)

// User struct
type User struct {
	ID   string `json:"id"`
	NAME string `json:"username"`
}

// GetUsers func
func GetUsers(rWriter http.ResponseWriter, req *http.Request) {
	db := connection.ConnectToDb()
	results, err := db.Query("SELECT id, username FROM users")
	if err != nil {
		panic(err.Error())
	}

	var id string
	var username string
	var users []User

	for results.Next() {
		err = results.Scan(&id, &username)
		if err != nil {
			panic(err.Error())
		}

		users = append(users, User{ID: id, NAME: username})
	}
	json.NewEncoder(rWriter).Encode(users)
}

// CreateUser func
func CreateUser(rWriter http.ResponseWriter, req *http.Request) {
	db := connection.ConnectToDb()
	insert, err := db.Query(" INSERT INTO users VALUES(5, 'MUTAKIN', 'password')")
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}

// GetUser func
func GetUser(rWriter http.ResponseWriter, req *http.Request) {
	db := connection.ConnectToDb()
	params := mux.Vars(req)
	userID := params["id"]
	fmt.Println(userID)
	results, err := db.Query("SELECT id, username FROM users WHERE id=?", userID)
	if err != nil {
		panic(err.Error())
	}

	var id string
	var username string
	var users []User

	for results.Next() {
		err = results.Scan(&id, &username)
		if err != nil {
			panic(err.Error())
		}

		users = append(users, User{ID: id, NAME: username})
	}
	json.NewEncoder(rWriter).Encode(users)
}
