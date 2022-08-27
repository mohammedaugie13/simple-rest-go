package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserAuth struct {
    UserId       string         `gorm:"column:userId;primary_key"`
	Email        string         `gorm:"column:email"`
	EmailVerif   bool           `gorm:"column:emailVerif"`
	PasswordHash string         `gorm:"column:passwordHash"`
	Status       int            `gorm:"column:status"`
	CreatedAt    time.Time      `gorm:"column:createdAt;autoCreateTime:nano"`
	UpdatedAt    time.Time      `gorm:"column:updatedAt;autoUpdateTime:nano"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deletedAt"`
}

func (UserAuth) TableName() string {
    return "userAuth"
}