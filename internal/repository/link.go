package repository

import (
	"URLProject/internal/entity"
	"URLProject/pkg/db"
)

type LinkRepository struct {
	database *db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{database: database}
}

func (lr *LinkRepository) Create(link *entity.Link) error {
	result := lr.database.DB.Create(link)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (lr *LinkRepository) Get(hash string) (*entity.Link, error) {
	var outputLink entity.Link
	result := lr.database.DB.First(&outputLink, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}
	return &outputLink, nil
}
