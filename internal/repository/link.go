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

func (lr *LinkRepository) Count() int64 {
	var count int64
	_ = lr.database.DB.
		Table("links").
		Where("deleted_at is NULL").
		Count(&count)
	return count
}

func (lr *LinkRepository) GetAll(limit, offset int) ([]entity.Link, error) {
	var links []entity.Link

	result := lr.database.DB.
		Table("links").
		Where("deleted_at is NULL").
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Scan(&links)
	if result.Error != nil {
		return nil, result.Error
	}
	return links, nil
}
