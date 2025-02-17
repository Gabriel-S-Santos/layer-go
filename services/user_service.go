package services

import (
	"errors"
	"strings"

	"layer-go/models"
	"layer-go/repositories"
)

// UserService define as operações de negócio para User.
type UserService interface {
	RegisterUser(name, email string) (*models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

// NewUserService cria uma nova instância de UserService.
func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) RegisterUser(name, email string) (*models.User, error) {
	// Validação: nome e email devem ter mais de 5 caracteres
	if len(name) <= 5 {
		return nil, errors.New("Nome deve ter mais de 5 caracteres")
	}
	if len(email) <= 5 || !strings.Contains(email, "@") {
		return nil, errors.New("E-mail inválido")
	}

	user := &models.User{
		Name:  name,
		Email: email,
	}

	// Persiste o usuário
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}
