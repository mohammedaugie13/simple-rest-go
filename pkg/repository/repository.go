package repository

import "gorm.io/gorm"

type Repository struct {
    UserAuthRepository IUserAuthRepository
}

func NewRepository(db *gorm.DB) *Repository {
    return &Repository{
        UserAuthRepository: NewUserAuthRepository(db),
    }
}