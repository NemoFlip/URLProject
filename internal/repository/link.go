package repository

import (
	"URLProject/internal/entity"
	"URLProject/pkg/db"
	"gorm.io/gorm/clause"
)

type LinkRepository struct {
	database *db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{database: database}
}

func (lr *LinkRepository) Create(link *entity.Link) error {
	result := lr.database.DB.Create(link)
	return result.Error
}

func (lr *LinkRepository) GetByHash(hash string) (*entity.Link, error) {
	var outputLink entity.Link
	result := lr.database.DB.First(&outputLink, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}
	return &outputLink, nil
}

func (lr *LinkRepository) Update(link *entity.Link) error {
	result := lr.database.DB.Clauses(clause.Returning{}).Updates(link)
	return result.Error
}

func (lr *LinkRepository) Delete(id uint) error {
	result := lr.database.DB.Delete(&entity.Link{}, id)
	return result.Error
}

func (lr *LinkRepository) GetByID(id uint) error {
	var outputLink entity.Link
	result := lr.database.DB.First(&outputLink, id)
	return result.Error
}
