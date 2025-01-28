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
