package service

import (
	"errors"
	"strings"
	"time"

	"github.com/Akakazkz/go-task-manager-api/internal/model"
	"github.com/Akakazkz/go-task-manager-api/internal/repository"
)

type TaskService interface {
	Create(task *model.Task) error
	ListByUserID(userID int64) ([]*model.Task, error)
	GetByID(id, userID int64) (*model.Task, error)
	Update(task *model.Task) error
	Delete(id, userID int64) error
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{
		repo: repo,
	}
}

func (s *taskService) Create(task *model.Task) error {
	if strings.TrimSpace(task.Title) == "" {
		return errors.New("title is required")
	}
	if task.DueDate != nil && task.DueDate.Before(time.Now()) {
		return errors.New("due date cannot be in the past")
	}
	return s.repo.Create(task)
}

func (s *taskService) ListByUserID(userID int64) ([]*model.Task, error) {
	return s.repo.ListByUserID(userID)
}

func (s *taskService) GetByID(id, userID int64) (*model.Task, error) {
	return s.repo.GetByID(id, userID)
}

func (s *taskService) Update(task *model.Task) error {
	if strings.TrimSpace(task.Title) == "" {
		return errors.New("title is required")
	}
	if task.DueDate != nil && task.DueDate.Before(time.Now()) {
		return errors.New("due date cannot be in the past")
	}
	return s.repo.Update(task)
}

func (s *taskService) Delete(id, userID int64) error{
	return s.repo.Delete(id, userID)
}