package main

import (
	"sync"
)

type StudentExamDetails struct {
	ExamScores map[int]float64 `json:"exam_scores"`
	Average float64 `json:"average"`
}

type ExamDetails struct{
	StudentExamScores map[string]float64 `json:"student_scores"`
	Average float64 `json:"average"`
}

type ScoreEvent struct{
	Exam int `json:"exam"`
	StudentID string `json:"studentId"`
	Score float64 `json:"score"`
}

var (
	studentScores = make(map[string]*StudentExamDetails)
	examScores = make(map[int]*ExamDetails)
	mu sync.Mutex
)