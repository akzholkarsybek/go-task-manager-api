package model

import "time"

type Task struct{
	ID int64 `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Completed bool `json:"completed"`
	DueDate *time.Time `json:"due_date"`
	UserID int64 `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}