package request

import (
	"log"
	"net/http"
	"restapi/service"

	"github.com/gorilla/mux"
)

// HandleRequest func
func HandleRequest() *mux.Router {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", rest.GetUsers).Methods("GET")
	myRouter.HandleFunc("/users/detail", rest.GetUser).Methods("GET")
	myRouter.HandleFunc("/users", rest.CreateUser).Methods("POST")
	myRouter.HandleFunc("/users", rest.UpdateUser).Methods("PUT")
	myRouter.HandleFunc("/users", rest.DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", myRouter))

	return myRouter
}
