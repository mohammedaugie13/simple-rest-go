package repository

import (
	"github.com/bitwyre/template-go/pkg/datastore/postgres/entity"
	"gorm.io/gorm"
)

type IUserAuthRepository interface {
	FindByEMail(email string) (entity.UserAuth, error)
}

type userAuthRepository struct {
	db *gorm.DB
}

func NewUserAuthRepository(db *gorm.DB) *userAuthRepository {
	return &userAuthRepository{db}
}

func (q *userAuthRepository) FindByEMail(email string) (entity.UserAuth, error) {
	var user entity.UserAuth

	err := q.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
