package handlers

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"golang-student-service/models"
	"log"
	"net/http"
	"time"
)

var validate *validator.Validate

func GetStudents(writer http.ResponseWriter, request *http.Request) {
}

func GetStudent(writer http.ResponseWriter, request *http.Request) {

}

// CreateStudent handles the creation of a new student
func CreateStudent(writer http.ResponseWriter, request *http.Request) {

	log.Print("CreateStudent called")

	var studentReq models.StudentRequest

	err := json.NewDecoder(request.Body).Decode(&studentReq)
	if err != nil {
		http.Error(writer, "Invalid request body"+err.Error(), http.StatusBadRequest)
		return
	}

	// Initialize the validator
	validate = validator.New()
	err = validate.Struct(studentReq)
	if err != nil {
		http.Error(writer, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	// save the student to the db

	studentRes := models.StudentResponse{
		FirstName: studentReq.FirstName,
		LastName:  studentReq.LastName,
		Age:       studentReq.Age,
		Grade:     studentReq.Grade,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(writer).Encode(studentRes)
	if err != nil {
		return
	}
}

func UpdateStudent(writer http.ResponseWriter, request *http.Request) {

}

func DeleteStudent(writer http.ResponseWriter, request *http.Request) {

}
