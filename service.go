package main

import (
	"database/sql"
	"net/http"
)
func EnableCors(w http.ResponseWriter, req *http.Request) {
	// Set CORS headers to allow requests from any frontend
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow requests from any origin
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT") // Allowed HTTP methods
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Allowed headers
	w.WriteHeader(http.StatusOK)
}

type Services struct{
	db *sql.DB
}

type Task struct{
	ID int
	Name string
}

func NewService(db *sql.DB) *Services {
	return &Services{db: db}
}

func (s *Services) AddTask(w http.ResponseWriter, req *http.Request){
	
	taskRequest := Task{}
	err := ReadFromRequestBody(req, &taskRequest)
	if err != nil {
		WriteErrToResponseBody(w,err)
		return
	}
	
	SQL := `INSERT INTO "tasks" (name) VALUES ($1) RETURNING id`
	err = s.db.QueryRow(SQL, taskRequest.Name).Scan(&taskRequest.ID)
	if err != nil {
		WriteErrToResponseBody(w,err)
		return
	}
	WriteToResponseBody(w, taskRequest)
}

func (s *Services) GetTasks(w http.ResponseWriter, req *http.Request){


	SQL := `SELECT id, name FROM tasks`
	rows, err := s.db.Query(SQL)
	if err != nil {
		WriteErrToResponseBody(w, err)
		return
	}
	defer rows.Close()

	tasks := []Task{}
	for rows.Next() {
		task := Task{}
		err = rows.Scan(&task.ID, &task.Name)
		if err != nil {
			WriteErrToResponseBody(w, err)
			return
		}
		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil{
		WriteErrToResponseBody(w, err)
		return
	}

	WriteToResponseBody(w, tasks)
}

func (s *Services) DeletTask(w http.ResponseWriter, req *http.Request){


	taskRequest := Task{}
	err := ReadFromRequestBody(req, &taskRequest)
	if err != nil{
		WriteErrToResponseBody(w, err)
		return
	}


	SQL := `DELETE FROM tasks WHERE id = $1`
	_, err = s.db.Exec(SQL,taskRequest.ID)
	if err != nil {
		WriteErrToResponseBody(w, err)
		return
	}
	s.GetTasks(w, req)

}

func (s *Services) EditTask(w http.ResponseWriter, req *http.Request){
	taskRequest := Task{}
	err := ReadFromRequestBody(req, &taskRequest)
	if err != nil {
		WriteErrToResponseBody(w, err)
		return
	}

	if taskRequest.ID == 0 {
		http.Error(w, "Task ID needed", http.StatusBadRequest)
		return
	}
	if taskRequest.Name == "" {
		http.Error(w, "Task name neded", http.StatusBadRequest)
		return
	}

	SQL := `UPDATE tasks SET name = $1 WHERE id = $2 RETURNING id, name`
	err = s.db.QueryRow(SQL, taskRequest.Name, taskRequest.ID).Scan(&taskRequest.ID, &taskRequest.Name)
	if err != nil {
		WriteErrToResponseBody(w, err)
		return
	}
	WriteToResponseBody(w, taskRequest)
}