package rest

import (
	"encoding/json"
	"net/http"
	"restapi/conn"
)

// User struct
type User struct {
	ID       string `json:"id"`
	NAME     string `json:"username"`
	PASSWORD string `json:"password"`
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
	id := req.FormValue("id")
	name := req.FormValue("name")
	password := req.FormValue("password")
	db := connection.ConnectToDb()
	insert, err := db.Prepare("INSERT INTO users (id, username, password) VALUES(?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	insert.Exec(id, name, password)

	defer db.Close()
}

// GetUser func
func GetUser(rWriter http.ResponseWriter, req *http.Request) {
	userID := req.URL.Query().Get("id")
	db := connection.ConnectToDb()
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
