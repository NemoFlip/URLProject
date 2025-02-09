package di

import "URLProject/internal/entity"

type IStatRepository interface {
	AddClick(linkId uint)
}

type IUserRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
