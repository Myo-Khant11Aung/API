package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func Listener() {
	resp, err := http.Get("http://live-test-scores.herokuapp.com/scores")
	if err != nil {
		log.Fatal("Failed to connect to SSE stream", err)
	}
	defer resp.Body.Close()

	reader := bufio.NewScanner(resp.Body)
	for reader.Scan() {
		line := reader.Text()
		if strings.HasPrefix(line, "data:") {
			var event ScoreEvent
			json.Unmarshal([]byte(line[6:]), &event)
			processData(event)
		}
	}
}

func processData(event ScoreEvent){
	mu.Lock()
	defer mu.Unlock()

	if _,exists := studentScores[event.StudentID]; !exists{
		studentScores[event.StudentID] = &StudentExamDetails{ExamScores: make(map[int]float64)}
	}
	studentScores[event.StudentID].ExamScores[event.Exam] = event.Score

	if _,exists := examScores[event.Exam]; !exists{
		examScores[event.Exam] = &ExamDetails{StudentExamScores: make(map[string]float64)}
	}
	examScores[event.Exam].StudentExamScores[event.StudentID] = event.Score

	total := 0.0
	count := 0

	for _, score := range studentScores[event.StudentID].ExamScores{
		total+= score
		count++
	}
	studentScores[event.StudentID].Average = total/ float64(count)

	examTotal := 0.0
	examCount := 0

	for _,score := range examScores[event.Exam].StudentExamScores{
		examTotal += score
		examCount++
	}
	examScores[event.Exam].Average = examTotal / float64(examCount)
}

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json")

	students := []string{}
	for studentID := range studentScores{
		students = append(students, studentID)
	}

	json.NewEncoder(w).Encode(students)
}

func GetStudentByID(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	studentID := vars["id"]

	student, exists := studentScores[studentID]
	if !exists {
		http.Error(w, "Student Not Found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(student)
}

func GetAllExams(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json")

	exams := []int{}
	for examNumber := range examScores {
		exams = append(exams, examNumber)
	}

	json.NewEncoder(w).Encode(exams)
}

func GetExamByNumber(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	examNumber, err := strconv.Atoi(vars["number"])

	if err != nil {
		http.Error(w, "Invalid exam number", http.StatusBadRequest)
		return
	}

	exam, exists := examScores[examNumber]
	if !exists {
		http.Error(w, "Exam not Found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(exam)

}