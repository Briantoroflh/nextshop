package users

import (
	"errors"
	"nextshop/entities"
)

type Service struct {
	Repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) CreateUser(user *entities.Users) error {
	if user.Email == "" {
		return errors.New("email are required")
	}
	return s.Repo.CreateUser(user)
}

func (s *Service) GetUserByID(id uint64) (*entities.Users, error) {
	return s.Repo.GetUserByID(id)
}

func (s *Service) UpdateUser(user *entities.Users) error {
	return s.Repo.UpdateUser(user)
}

func (s *Service) DeleteUser(id uint64) error {
	return s.Repo.DeleteUser(id)
}

func (s *Service) GetAllUsers() (*entities.Users, error) {
	return s.Repo.GetAllUsers()
}

func (s *Service) GetUserByRole(role *entities.Users) (*entities.Users, error) {
	return s.Repo.GetUserByRole(role)
}