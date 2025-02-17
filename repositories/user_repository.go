package repositories

import (
	"layer-go/database"
	"layer-go/models"
)

// UserRepository define as operações de persistência para User.
type UserRepository interface {
	Create(user *models.User) error
}

// userRepository é a implementação de UserRepository usando Gorm.
type userRepository struct{}

// NewUserRepository cria uma nova instância de userRepository.
func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(user *models.User) error {
	return database.DB.Create(user).Error
}
