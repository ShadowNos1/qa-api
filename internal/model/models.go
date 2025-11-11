package model

import "time"

var ErrNotFound = &NotFoundError{}

type NotFoundError struct{}

func (e *NotFoundError) Error() string {
    return "not found"
}

type Question struct {
    ID        uint      `json:"id"`
    Text      string    `json:"text"`
    CreatedAt time.Time `json:"created_at"`
}

type Answer struct {
    ID         uint      `json:"id"`
    QuestionID uint      `json:"question_id"`
    UserID     string    `json:"user_id"`
    Text       string    `json:"text"`
    CreatedAt  time.Time `json:"created_at"`
}
