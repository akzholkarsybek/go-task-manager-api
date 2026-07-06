package service

import (
	"strings"
	"time"

	"github.com/Akakazkz/go-task-manager-api/internal/model"
	"github.com/Akakazkz/go-task-manager-api/internal/repository"
	"github.com/Akakazkz/go-task-manager-api/internal/auth"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) List() ([]*model.User, error) {
	return s.repo.List()
}

func (s *userService) Login(email, password string) (string, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	if err := comparePassword(user.Password, password); err != nil {
		return "", ErrInvalidCredentials
	}

	token, err := auth.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) Create(email, password string) (*model.User, error) {
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)

	if email == "" || password == "" {
		return nil, ErrInvalidInput
	}

	if s.repo.ExistsByEmail(email) {
		return nil, ErrUserExists
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		Email:     email,
		Password:  hashedPassword,
		Role:      model.RoleUser,
		CreatedAt: time.Now(),
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}
