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
	myRouter.HandleFunc("/users/{id}", rest.GetUser).Methods("GET")
	myRouter.HandleFunc("/users", rest.CreateUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))

	return myRouter
}
