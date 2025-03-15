
# SSE Score Tracker API

This is a Golang-based API that listens to Server-Sent Events (SSE) from a live test score stream, processes the data, and provides RESTful APIs to access student scores and exam results.


## Features

- Consumes SSE Data: Listens to real-time score events.
- In-Memory Storage: Stores processed results in maps (no database required).
- REST API: Exposes endpoints to retrieve student scores and exam details.
- Gorilla Mux: Used for efficient request routing.
- Concurrent & Thread-Safe: Uses goroutines and mutex locks to handle concurrent updates.

##  Libraries Used
- net/http → Handles HTTP server and API requests
- encoding/json → JSON parsing & encoding
- bufio & strings → Reads and processes SSE stream
- sync.Mutex → Prevents race conditions in concurrent updates
- github.com/gorilla/mux → Router for handling API requests







## How to run the API

Open terminal or bash and go to the project directory on your machine.

- Run the command below to install dependencies
```bash
go mod tidy
```

- Run the command below to start the program
```bash
go run .
```

The server should now be running at
```bash
http://localhost:8080
```

## Running Tests

Here’s how you can test the API using curl:

- Get All Students
Returns a list of students who received at least one test score.

```bash
curl -X GET http://localhost:8080/students
```

- Get a Specific Student’s Scores
Returns all test scores and the average score for a given student.

```bash
curl -X GET http://localhost:8080/students/{id}
```

- Get All Exams
Returns a list of all recorded exam numbers.

```bash
curl -X GET http://localhost:8080/exams
```

- Get Exam Results & Average Score
Returns all student scores for a specific exam and its overall average.

```bash
curl -X GET http://localhost:8080/exams/{number}
```

## Author
[@Myo-Khant11Aung](https://github.com/Myo-Khant11Aung)

