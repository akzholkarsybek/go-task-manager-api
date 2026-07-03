package repository

import "github.com/Akakazkz/go-task-manager-api/internal/model"

type TaskRepository interface{
	Create(task *model.Task) error
	GetByID(id, userID int64) (*model.Task, error)
	ListByUserID(userID int64) ([]*model.Task, error)
	Update(task *model.Task) error
	Delete(id, userID int64) error
}