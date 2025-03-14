package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func main() {
	router := mux.NewRouter()

	router.HandleFunc("/students", GetAllStudents).Methods("GET")
	router.HandleFunc("/students/{id}", GetStudentByID).Methods("GET")
	router.HandleFunc("/exams", GetAllExams).Methods("GET")
	router.HandleFunc("/exams/{number}", GetExamByNumber).Methods("GET")

	go Listener()

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}