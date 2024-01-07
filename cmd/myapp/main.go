package main

import (
	"github.com/gorilla/mux"
	"golang-student-service/db"
	"golang-student-service/handlers"
	"log"
	"net/http"
)

func main() {

	_, err := db.ConnectDB()

	if err != nil {
		log.Fatalf("Could not connect to db: %s\n", err.Error())
	}

	r := mux.NewRouter()
	r.HandleFunc("/students", handlers.GetStudents).Methods("GET")
	r.HandleFunc("/students/{id}", handlers.GetStudent).Methods("GET")
	r.HandleFunc("/students", handlers.CreateStudent).Methods("POST")
	r.HandleFunc("/students/{id}", handlers.UpdateStudent).Methods("PUT")
	r.HandleFunc("/students/{id}", handlers.DeleteStudent).Methods("DELETE")
	err = http.ListenAndServe(":8000", r)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
