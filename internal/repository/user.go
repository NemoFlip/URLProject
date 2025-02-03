package repository

import (
	"URLProject/internal/entity"
	"URLProject/pkg/db"
)

type UserRepository struct {
	database *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{database: database }
}

func (ur *UserRepository) Create(user *entity.User) error {
	result := ur.database.DB.Create(user)
	return result.Error
}

func (ur *UserRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	result := ur.database.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
