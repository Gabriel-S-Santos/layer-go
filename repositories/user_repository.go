package repositories

import (
	"layer-go/database"
	"layer-go/models"
)

type UserRepository interface {
	Create(user *models.User) error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(user *models.User) error {
	return database.DB.Create(user).Error
}
